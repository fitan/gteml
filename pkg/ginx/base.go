package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type GinX struct {
	*gin.Context
	bindReq interface{}
	bindRes interface{}
	bindErr error

	entryMiddleware *[]types.Middleware

	handlerMiddleware *[]types.Middleware
}

func (g *GinX) SetEntryMid(m *[]types.Middleware) {
	g.entryMiddleware = m
}

func (g *GinX) SetHandlerMid(m *[]types.Middleware) {
	g.handlerMiddleware = m
}

func (g *GinX) BindReq() interface{} {
	return g.bindReq
}

func (g *GinX) BindRes() interface{} {
	return g.bindRes
}

func (g *GinX) BindErr() error {
	return g.bindErr
}

func (g *GinX) SetBindReq(i interface{}) {
	g.bindReq = i
}

func (g *GinX) SetBindRes(i interface{}) {
	g.bindRes = i
}

func (g *GinX) Reset() {
	g.bindReq = nil
	g.bindErr = nil
	g.bindRes = nil
	g.handlerMiddleware = nil
	g.Context = nil
}

func (g *GinX) SetBindErr(err error) {
	if err == nil {
		return
	}
	if g.bindErr == nil {
		g.bindErr = err
		return
	}

	g.bindErr = errors.Wrapf(g.bindErr, "%s :", err.Error())
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
		core.Log.Sync()
		core.Pool.ReUse(core)
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
	g.bindReq = data
	g.SetBindErr(err)
}

func (g *GinX) setBindFn(data interface{}, err error) {
	g.bindRes = data
	g.SetBindErr(err)
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
