package core

import (
	"github.com/fitan/magic/pkg/types"
)

type PoolRegister struct {
}

func (p *PoolRegister) With(o ...types.Option) types.Register {
	return p
}

func (p *PoolRegister) Reload(c *types.Core) {
}

func (p *PoolRegister) Set(c *types.Core) {
	c.Pool = GetCorePool()
}

func (p *PoolRegister) Unset(c *types.Core) {
}
