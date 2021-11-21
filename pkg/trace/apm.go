package trace

import (
	"context"
	"go.elastic.co/apm"
	"go.opentelemetry.io/otel/trace"
)

func NewApmTrace() {

}

type ApmTrace struct {
	isOpen bool
	ctx    context.Context
	spans  []*apm.Span
}

func (a *ApmTrace) IsOpen() bool {
	return a.isOpen
}

func (a *ApmTrace) UnSet() {
	a.ctx = nil
	a.spans = a.spans[0:0]
}

func (a *ApmTrace) Tp() *trace.TracerProvider {
	return nil
}

func (a *ApmTrace) SpanCtx(name string) context.Context {
	span, nextCtx := apm.StartSpan(a.ctx, name, "method")
	a.spans = append(a.spans, span)
	a.ctx = nextCtx
	return a.ctx
}

func (t *ApmTrace) SetCtx(ctx context.Context) {
	t.ctx = ctx
}

func (a *ApmTrace) Ctx() context.Context {
	return a.ctx
}

func (t *ApmTrace) End() {
	for _, s := range t.spans {
		if !s.IsExitSpan() {
			s.End()
		}
	}
}
