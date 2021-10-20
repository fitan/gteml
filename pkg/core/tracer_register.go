package core

import "github.com/fitan/gteml/pkg/trace"

type Trace struct {
}

func (t Trace) Reload(c *Context) {
	panic("implement me")
}

func (t Trace) With(o ...Option) Register {
	panic("implement me")
}

func (t Trace) Set(c *Context) {
	c.Tracer = trace.NewTrace(trace.GetTp(), c.Config.Trace.Open)
}

func (t Trace) Unset(c *Context) {
	c.Tracer = nil
}
