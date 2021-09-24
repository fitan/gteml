package core

import (
	"context"
	"github.com/fitan/gteml/pkg/log"
)

type Register interface {
	Set(c *Context)
	Unset(c *Context)
}

type Option func(c *Context)

func WithInitRegister(rs ...Register) Option {
	return func(c *Context) {
		for _, r := range rs {
			c.Register(r)
		}
	}
}

type Context struct {
	Log *CoreLog

	objRegister []Register

	//Log *log.Xlog

	TraceLog *log.TraceLog

	GinX *GinX

	Ctx context.Context

	releaseFn func(x interface{})
}

func (c *Context) Register(register Register) {
	c.objRegister = append(c.objRegister, register)
}

func (c *Context) Init() {
	for _, m := range c.objRegister {
		m.Set(c)
	}
}

func (c *Context) Release() {
	for _, m := range c.objRegister {
		m.Unset(c)
	}

	c.releaseFn(c)
}
