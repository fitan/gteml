package context

import (
	"github.com/fitan/magic/pkg/trace"
	"github.com/fitan/magic/pkg/types"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

type Trace struct {
	tp *oteltrace.TracerProvider
}

func (t *Trace) getTp(c *types.Context) *oteltrace.TracerProvider {
	if t.tp == nil {
		tp, err := trace.TracerProvider(c.Config.App.Name, c.Config.Trace.TracerProviderAddr)
		if err != nil {
			log.Println(err)
			return nil
		}
		t.tp = tp
	}
	return t.tp
}

func (t *Trace) Reload(c *types.Context) {
	t.tp = nil
}

func (t *Trace) With(o ...types.Option) types.Register {
	return nil
}

func (t *Trace) Set(c *types.Context) {
	c.Tracer = trace.NewTrace(t.getTp(c), c.Config.Trace.Open)
}

func (t *Trace) Unset(c *types.Context) {
	c.Tracer.UnSet()
}
