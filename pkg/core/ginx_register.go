package core

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
)

type GinXRegister struct {
	EntryMid []types.Middleware
}

func (g *GinXRegister) Reload(c *types.Core) {
}

func (g *GinXRegister) With(o ...types.Option) types.Register {
	return g
}

func (g *GinXRegister) Set(c *types.Core) {
	c.GinX = ginx.NewGin(ginx.WithEntryMid(&g.EntryMid))
}

func (g *GinXRegister) Unset(c *types.Core) {
	c.GinX.Reset()
}
