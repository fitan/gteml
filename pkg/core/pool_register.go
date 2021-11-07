package core

import (
	"github.com/fitan/magic/pkg/types"
	"sync"
)

var CtxPool *types.CorePool

func GetCorePool() *types.CorePool {
	if CtxPool == nil {
		CtxPool = &types.CorePool{}
		CtxPool.P = sync.Pool{New: NewObjFn(CtxPool)}
	}
	return CtxPool
}

func NewObjFn(p *types.CorePool) func() interface{} {
	return func() interface{} {
		c := &types.Core{}
		p.Set(c)
		c.LocalVersion = c.Version.Version()
		return c
	}
}

type PoolReg struct {
}

func (p *PoolReg) With(o ...types.Option) types.Register {
	return p
}

func (p *PoolReg) Reload(c *types.Core) {
}

func (p *PoolReg) Set(c *types.Core) {
	c.Pool = GetCorePool()
}

func (p *PoolReg) Unset(c *types.Core) {
}
