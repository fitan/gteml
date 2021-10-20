package core

import (
	"github.com/fitan/gteml/pkg/types"
)

type Register interface {
	With(o ...Option) Register
	Reload(c *Context)
	Set(c *Context)
	Unset(c *Context)
}

type Option func(c *Context)

type Context struct {
	Config *types.MyConf

	*CoreLog

	Log types.Logger

	Tracer types.Tracer

	GinX types.GinXer

	Storage types.Storage

	Cache types.Cache

	Apis Apis

	Version      types.Version
	localVersion int

	releaseFn func(x interface{})
}

func (c *Context) Init() *Context {
	for _, r := range registerList {
		r.Set(c)
	}

	return c
}

func (c *Context) SetLocalVersion() {
	c.localVersion = c.Version.Version()
}

func (c *Context) Release() {
	c.Tracer.End()
	for _, r := range registerList {
		r.Unset(c)
	}

	// 如果配置文件reload 那么对象不放回pool中
	if c.localVersion < c.Version.Version() {
		return
	}

	PutCore(c)
}
