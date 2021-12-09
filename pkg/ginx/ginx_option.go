package ginx

import (
	"github.com/fitan/magic/pkg/types"
)

const FnName = "fn_name"

func WithEntryMid(m *[]types.Middleware) GinXOption {
	return func(g *GinX) {
		g.SetEntryMid(m)
	}
}

type GinXHandlerOption func(c *types.Core) error

// gin value 设置key
func WithHandlerName(name string) GinXHandlerOption {
	return func(c *types.Core) error {
		c.GinX.GinCtx().Set(FnName, name)
		return nil
	}
}

func WithHandlerMid(mids ...types.Middleware) GinXHandlerOption {
	ms := make([]types.Middleware, 0, len(mids))
	ms = append(ms, mids...)
	return func(core *types.Core) error {
		core.GinX.SetHandlerMid(&ms)
		return nil
	}
}
