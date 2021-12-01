package core

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
)

type ginXRegister struct {
	EntryMid []types.Middleware
}

func (g *ginXRegister) Reload(c *types.Core) {
}

func (g *ginXRegister) With(o ...types.Option) types.Register {
	return g
}

func (g *ginXRegister) Set(c *types.Core) {
	c.GinX = ginx.NewGin(ginx.WithEntryMid(&g.EntryMid))
}

func (g *ginXRegister) Unset(c *types.Core) {
	c.GinX.Reset()
}
