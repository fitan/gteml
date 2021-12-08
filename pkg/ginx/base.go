package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type GinX struct {
	*gin.Context
	request  interface{}
	response interface{}
	errors   []error

	entryMiddleware *[]types.Middleware

	handlerMiddleware *[]types.Middleware
}

func (g *GinX) SetEntryMid(m *[]types.Middleware) {
	g.entryMiddleware = m
}

func (g *GinX) SetHandlerMid(m *[]types.Middleware) {
	g.handlerMiddleware = m
}

func (g *GinX) Request() interface{} {
	return g.request
}

func (g *GinX) Response() interface{} {
	return g.response
}

func (g *GinX) LastError() error {
	if len(g.errors) == 0 {
		return nil
	}
	return g.errors[len(g.errors)-1]
}

func (g *GinX) Errors() []error {
	return g.errors
}

func (g *GinX) SetRequest(i interface{}) {
	g.request = i
}

func (g *GinX) SetResponse(i interface{}) {
	g.response = i
}

func (g *GinX) Reset() {
	g.request = nil
	g.errors = g.errors[:0]
	g.response = nil
	g.handlerMiddleware = nil
	g.Context = nil
}

func (g *GinX) SetError(err error) {
	if err == nil {
		return
	}
	g.errors = append(g.errors, err)
	return
}

func (g *GinX) BindTransfer(core *types.Core, i types.GinXBinder) {
	defer func() {
		if g.entryMiddleware != nil {
			for _, fn := range *g.entryMiddleware {
				fn.Forever(core)
			}
		}

		if g.handlerMiddleware != nil {
			for _, fn := range *g.handlerMiddleware {
				fn.Forever(core)
			}
		}
		//core.Log.Sync()
		//core.Pool.ReUse(core)
	}()

	if g.entryMiddleware != nil {
		for _, fn := range *g.entryMiddleware {
			if !fn.BindValBefor(core) {
				return
			}
		}
	}

	if g.handlerMiddleware != nil {
		for _, fn := range *g.handlerMiddleware {
			if !fn.BindValBefor(core) {
				return
			}
		}
	}

	g.setBindVal(i.BindVal(core))

	if g.entryMiddleware != nil {
		for _, fn := range *g.entryMiddleware {
			if !fn.BindValAfter(core) {
				return
			}
		}
	}

	if g.handlerMiddleware != nil {
		for _, fn := range *g.handlerMiddleware {
			if !fn.BindValAfter(core) {
				return
			}
		}
	}

	g.setBindFn(i.BindFn(core))

	if g.entryMiddleware != nil {
		for _, fn := range *g.entryMiddleware {
			if !fn.BindFnAfter(core) {
				return
			}
		}
	}

	if g.handlerMiddleware != nil {
		for _, fn := range *g.handlerMiddleware {
			if !fn.BindFnAfter(core) {
				return
			}
		}
	}
}

func (g *GinX) SetGinCtx(c *gin.Context) {
	g.Context = c
}

func (g *GinX) GinCtx() *gin.Context {
	return g.Context
}

func (g *GinX) setBindVal(data interface{}, err error) {
	g.request = data
	g.SetError(errors.WithMessage(err, "bind val error"))
}

func (g *GinX) setBindFn(data interface{}, err error) {
	g.response = data
	g.SetError(errors.WithMessage(err, "bind fn error"))
}

type GinXOption func(g *GinX)

func NewGin(fs ...GinXOption) *GinX {
	g := new(GinX)
	for _, f := range fs {
		f(g)
	}
	return g
}

type GinXHandlerOption func(c *types.Core) error

// gin value 设置key
func WithHandlerName(name string) GinXHandlerOption {
	return func(c *types.Core) error {
		c.GinX.GinCtx().Set(FnName, name)
		return nil
	}
}
