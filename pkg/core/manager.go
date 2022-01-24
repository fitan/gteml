package core

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
	"sync"
)

//type ContextPool struct {
//	p            sync.Pool
//	registerList []Register
//}
//
//func ConfReloadHook() {
//}
//
//var registerList []Register
//
//func InsetRegister(os ...Register) {
//	registerList = append(registerList, os...)
//}

var ConfReg *ConfRegister
var PromReg *PromRegister

func NewCore() {
	ConfReg = NewConfReg()
	PromReg = NewPromRegister()
	GetCorePool().RegisterList([]types.Register{
		ConfReg,
		&TraceRegister{},
		&LogRegister{},
		&GinXRegister{
			EntryMid: []types.Middleware{
				&ginx.ResultWrapMid{},
				&ginx.TraceMid{},
			},
		},
		&DaoRegister{},
		PromReg,
		NewServiceRegister(),
		&ApisRegister{},
		&VersionRegister{},
		&PoolRegister{},
	})
}

type CorePool struct {
	P            sync.Pool
	registerList []types.Register
}

func (c *CorePool) RegisterList(l []types.Register) {
	c.registerList = l
}

func (c *CorePool) Set(ctx *types.Core) {
	for _, v := range c.registerList {
		v.Set(ctx)
	}
}

func (c *CorePool) Unset(ctx *types.Core) {
	for _, v := range c.registerList {
		v.Unset(ctx)
	}
}

func (c *CorePool) Reload() {
	ctx := c.GetObj()
	for _, v := range c.registerList {
		v.Reload(ctx)
	}
}

// 从pool获取对象后进行初始化
func (c *CorePool) GetInit() {
	// Todo 获取pool后的初始化
}

func (c *CorePool) ReUse(ctx *types.Core) {
	// tracer收尾 防止有的trace 没有end
	ctx.Tracer.End()

	c.Unset(ctx)

	// 如果配置文件reload 那么对象不放回pool中
	if ctx.LocalVersion != ctx.Version.Version() {
		PromReg.Get().CorePool("!Put:Version")
		return
	}

	PromReg.Get().CorePool("Put")
	c.P.Put(ctx)
}

func (c *CorePool) GetObj() *types.Core {
	for {
		context := c.P.Get().(*types.Core)
		if context.LocalVersion != context.Version.Version() {
			PromReg.Get().CorePool("!Get:Version")
			continue
		}
		PromReg.Get().CorePool("Get")
		return context
	}
}

var CtxPool *CorePool

func GetCorePool() *CorePool {
	if CtxPool == nil {
		CtxPool = &CorePool{}
		CtxPool.P = sync.Pool{New: NewObjFn(CtxPool)}
	}
	return CtxPool
}

func NewObjFn(p *CorePool) func() interface{} {
	return func() interface{} {
		c := &types.Core{}
		p.Set(c)
		c.LocalVersion = c.Version.Version()
		PromReg.Get().CorePool("Create")
		return c
	}
}
