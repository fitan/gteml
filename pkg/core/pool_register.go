package core

import (
	"github.com/fitan/magic/pkg/types"
)

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
