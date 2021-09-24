package core

import (
	"github.com/gin-gonic/gin"
)

type GinXTransfer interface {
	BindVal(c *Context) (interface{}, error)
	BindFn(c *Context) (interface{}, error)
}

type GinX struct {
	*gin.Context
	xkeys    map[string]string
	bindData struct {
		// 入参
		val interface{}
		// 出参
		res interface{}
		err error
	}
	resultWrap []Option
}

func (g *GinX) BindTransfer(core *Context, ctx *gin.Context, i GinXTransfer) {
	defer g.result(core)
	g.setCtx(ctx)

	g.setBindVal(i.BindVal(core))
	if g.checkErr() {
		return
	}

	g.setBindVal(i.BindFn(core))
	if g.checkErr() {
		return
	}

	core.Release()
}

func (g *GinX) checkErr() bool {
	if g.bindData.err != nil {
		return true
	}
	return false
}

func (g *GinX) setCtx(c *gin.Context) {
	g.Context = c
}

func (g *GinX) Ctx() *gin.Context {
	return g.Context
}

func (g *GinX) setBindVal(data interface{}, err error) {
	g.bindData.val = data
	g.bindData.err = err
}

func (g *GinX) setBindFn(data interface{}, err error) {
	g.bindData.res = data
	g.bindData.err = err
}

func (g *GinX) result(c *Context) {
	for _, r := range g.resultWrap {
		r(c)
	}
}

type GinOption func(g *GinX)

func NewGin(fs ...GinOption) *GinX {
	g := &GinX{}
	for _, f := range fs {
		f(g)
	}
	return g
}

type ginXRegister struct {
}

func (g *ginXRegister) Set(c *Context) {
	c.GinX = NewGin(WithWrap(GinResultWrap, GinTraceWrap))
}

func (g *ginXRegister) Unset(c *Context) {
	c.GinX = nil
}

func GinXHandlerRegister(t func() GinXTransfer, o ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		core := GetCore().(*Context)
		core.GinX.BindTransfer(core, c, t())
	}
}
