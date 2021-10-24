package core

import (
	"github.com/fitan/gteml/pkg/types"
	"github.com/gin-gonic/gin"
)

type GinX struct {
	*gin.Context
	bindReq    interface{}
	bindRes    interface{}
	bindErr    error
	resultWrap []types.Option
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

func (g *GinX) SetBindErr(err error) {
	g.bindErr = err
}

func NewGinX() types.GinXer {
	return &GinX{}
}

func (g *GinX) BindTransfer(core *types.Context, i types.GinXBinder) {
	defer core.Pool.ReUse(core)
	defer g.Result(core)
	defer core.Log.Sync()
	if g.checkErr() {
		return
	}

	g.setBindVal(i.BindVal(core))
	if g.checkErr() {
		return
	}

	g.setBindFn(i.BindFn(core))
}

func (g *GinX) checkErr() bool {
	if g.BindErr() != nil {
		return true
	}
	return false
}

func (g *GinX) SetGinCtx(c *gin.Context) {
	g.Context = c
}

func (g *GinX) GinCtx() *gin.Context {
	return g.Context
}

func (g *GinX) setBindVal(data interface{}, err error) {
	g.bindReq = data
	g.bindErr = err
}

func (g *GinX) setBindFn(data interface{}, err error) {
	g.bindRes = data
	g.bindErr = err
}

func (g *GinX) Result(c *types.Context) {
	for _, r := range g.resultWrap {
		r(c)
	}
}

type GinOption func(g *GinX)

func NewGin(fs ...GinOption) *GinX {
	g := new(GinX)
	for _, f := range fs {
		f(g)
	}
	return g
}

type ginXRegister struct {
	options []types.Option
}

func (g *ginXRegister) Reload(c *types.Context) {
}

func (g *ginXRegister) With(o ...types.Option) types.Register {
	g.options = append(make([]types.Option, 0, len(o)), o...)
	return g
}

func (g *ginXRegister) Set(c *types.Context) {
	c.GinX = NewGin(WithWrap(GinXResultWrap, GinXTraceWrap))
}

func (g *ginXRegister) Unset(c *types.Context) {
	c.GinX.SetBindReq(nil)
	c.GinX.SetBindRes(nil)
	c.GinX.SetBindErr(nil)
	c.GinX.SetGinCtx(nil)
}

type GinXHandlerOption func(c *types.Context) error

func GinXHandlerRegister(i gin.IRouter, transfer types.GinXTransfer, o ...GinXHandlerOption) {
	i.Handle(transfer.Method(), transfer.Url(), func(c *gin.Context) {
		core := GetCtxPool().GetObj()
		//gin的request ctx放到trace里
		//core.SetCtx(c.Request.Context())
		// core包裹gin context
		core.GinX.SetGinCtx(c)
		core.Tracer.SetCtx(c.Request.Context())
		// 加载中间件option
		for _, f := range o {
			err := f(core)
			if err != nil {
				core.GinX.SetBindErr(err)
				break
			}
		}

		if core.Tracer.IsOpen() {

			if core.CoreLog.IsOpenTrace() {
				// 设置tracelog
				core.Log = core.CoreLog.TraceLog(core.GinX.GinCtx().GetString(_FnName))
				// 如果打开trace则end
				defer core.Log.Sync()
			} else {
				core.Log = core.CoreLog.Log()
			}
		} else {
			// 普通log
			core.Log = core.CoreLog.Log()
		}

		core.GinX.BindTransfer(core, transfer.Binder())
	})
}

// gin value 设置key
func WithHandlerName(name string) GinXHandlerOption {
	return func(c *types.Context) error {
		c.GinX.GinCtx().Set(_FnName, name)
		return nil
	}
}
