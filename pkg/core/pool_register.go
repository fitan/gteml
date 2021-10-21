package core

import (
	"github.com/fitan/gteml/pkg/types"
	"sync"
)

var CtxPool *types.CtxPool

func GetCtxPool() *types.CtxPool {
	if CtxPool == nil {
		CtxPool = &types.CtxPool{P: sync.Pool{New: NewObjFn(CtxPool)}}
		return CtxPool
	}
	return CtxPool
}

func NewObjFn(p *types.CtxPool) func() interface{} {
	return func() interface{} {
		c := &types.Context{}
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

func (p *PoolReg) Reload(c *types.Context) {
}

func (p *PoolReg) Set(c *types.Context) {
	c.Pool = GetCtxPool()
}

func (p *PoolReg) Unset(c *types.Context) {
}
