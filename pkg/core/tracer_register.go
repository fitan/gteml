package core

import (
	"github.com/fitan/magic/pkg/trace"
	"github.com/fitan/magic/pkg/types"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

type Trace struct {
	tp *oteltrace.TracerProvider
}

func (t *Trace) getTp(c *types.Core) *oteltrace.TracerProvider {
	if t.tp == nil {
		tp, err := trace.TracerProvider(ConfReg.Confer.GetMyConf().App.Name, ConfReg.Confer.GetMyConf().Trace.TracerProviderAddr)
		if err != nil {
			log.Println(err)
			return nil
		}
		t.tp = tp
	}
	return t.tp
}

func (t *Trace) Reload(c *types.Core) {
	t.tp = nil
}

func (t *Trace) With(o ...types.Option) types.Register {
	return nil
}

func (t *Trace) Set(c *types.Core) {
	c.Tracer = trace.NewTrace(t.getTp(c), ConfReg.Confer.GetMyConf().Trace.Open)
}

func (t *Trace) Unset(c *types.Core) {
	c.Tracer.UnSet()
}
