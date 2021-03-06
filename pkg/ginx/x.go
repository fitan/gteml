package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"runtime"
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

		err := recover()
		if err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log := core.GetCoreLog().TraceLog("pkg.ginX.wrapMid.recover")
			log.Error("panic", zap.Any("err", err), zap.String("stack", string(buf)))
			log.Sync()

			panic(err)
			core.GinX.SetError(errors.New("系统错误，请联系管理员"))
		}

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
		//core.log.Sync()
		//core.Pool.ReUse(core)
	}()

	if g.entryMiddleware != nil {
		for _, fn := range *g.entryMiddleware {
			if !fn.BindValBefore(core) {
				return
			}
		}
	}

	if g.handlerMiddleware != nil {
		for _, fn := range *g.handlerMiddleware {
			if !fn.BindValBefore(core) {
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
	g.SetError(errors.Wrap(err, "bind val error"))
}

func (g *GinX) setBindFn(data interface{}, err error) {
	g.response = data
	g.SetError(errors.Wrap(err, "bind fn error"))
}

type GinXOption func(g *GinX)

func NewGin(fs ...GinXOption) *GinX {
	g := new(GinX)
	for _, f := range fs {
		f(g)
	}
	return g
}
