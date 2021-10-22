package core

import (
	"github.com/fitan/gteml/pkg/trace"
	"github.com/fitan/gteml/pkg/types"
)

type Trace struct {
}

func (t *Trace) Reload(c *types.Context) {
}

func (t *Trace) With(o ...types.Option) types.Register {
	return nil
}

func (t *Trace) Set(c *types.Context) {
	c.Tracer = trace.NewTrace(trace.GetTp(), c.Config.Trace.Open)
}

func (t *Trace) Unset(c *types.Context) {
}
