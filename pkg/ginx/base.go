package ginx

import (
	"github.com/fitan/magic/pkg/types"
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

type GinXHandlerOption func(c *types.Context) error

// gin value 设置key
func WithHandlerName(name string) GinXHandlerOption {
	return func(c *types.Context) error {
		c.GinX.GinCtx().Set(FnName, name)
		return nil
	}
}
