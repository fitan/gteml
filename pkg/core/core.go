package core

import (
	"github.com/fitan/gteml/pkg/common"
)

type Register interface {
	With(o ...Option) Register
	Set(c *Context)
	Unset(c *Context)
}

type Option func(c *Context)

type Context struct {
	*CoreLog

	Log common.Logger

	common.Tracer

	GinX *GinX

	Store interface{}

	Apis

	releaseFn func(x interface{})
}

func (c *Context) Init() *Context {
	for _, r := range registerList {
		r.Set(c)
	}

	return c
}

func (c *Context) Release() {
	c.Tracer.End()
	for _, r := range registerList {
		r.Unset(c)
	}

	PutCore(c)
}
