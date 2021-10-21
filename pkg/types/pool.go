package types

import (
	"sync"
)

type CtxPool struct {
	P            sync.Pool
	registerList []Register
}

func (c *CtxPool) RegisterList(l []Register) {
	c.registerList = l
}

func (c *CtxPool) Set(ctx *Context) {
	for _, v := range c.registerList {
		v.Set(ctx)
	}
}

func (c *CtxPool) Unset(ctx *Context) {
	for _, v := range c.registerList {
		v.Unset(ctx)
	}
}

func (c *CtxPool) Reload(ctx *Context) {
	for _, v := range c.registerList {
		v.Reload(ctx)
	}
}

// 从pool获取对象后进行初始化
func (c *CtxPool) GetInit() {
	// Todo 获取pool后的初始化
}

func (c *CtxPool) ReUse(ctx *Context) {
	// tracer收尾 防止有的trace 没有end
	ctx.Tracer.End()

	c.Unset(ctx)

	// 如果配置文件reload 那么对象不放回pool中
	if ctx.LocalVersion < ctx.Version.Version() {
		return
	}

	c.P.Put(ctx)
}

func (c *CtxPool) GetObj() *Context {
	return c.P.Get().(*Context)

}

var registerList []Register

type Register interface {
	With(o ...Option) Register
	Reload(c *Context)
	Set(c *Context)
	Unset(c *Context)
}
type Pooler interface {
	RegisterList(l []Register)
	Set(ctx *Context)
	Unset(ctx *Context)
	Reload(ctx *Context)
	GetInit()
	ReUse(ctx *Context)
	GetObj() *Context
}
