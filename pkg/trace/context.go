package trace

import (
	"context"
	"go.elastic.co/apm"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	trace2 "go.opentelemetry.io/otel/trace"
	"sync"
)

func NewTrace(tp *trace.TracerProvider, open bool) *Trace {
	t := new(Trace)
	t.isOpen = open
	t.tp = tp
	return t
}

type Trace struct {
	tp       *trace.TracerProvider
	ctx      context.Context
	m        sync.Mutex
	isOpen   bool
	spans    []trace2.Span
	apmSpans []*apm.Span
}

func (t *Trace) Tp() *trace.TracerProvider {
	return t.tp
}

func (t *Trace) UnSet() {
	t.ctx = nil
	t.spans = t.spans[0:0]
}

func (t *Trace) IsOpen() bool {
	t.m.Lock()
	defer t.m.Unlock()
	return t.isOpen
}

func (t *Trace) SetCtx(ctx context.Context) {
	t.m.Lock()
	defer t.m.Unlock()
	t.ctx = ctx
}

func (t *Trace) Ctx() context.Context {
	t.m.Lock()
	defer t.m.Unlock()
	if t.ctx == nil {
		return context.Background()
	}
	return t.ctx
}

func (t *Trace) SpanCtx(name string) context.Context {
	t.m.Lock()
	defer t.m.Unlock()
	spanCtx, span := t.tp.Tracer(name).Start(t.ctx, name)
	t.spans = append(t.spans, span)
	t.ctx = spanCtx
	return spanCtx
}

//func (t *Trace) ApmSpanCtx(name string, spanType string) (context.Context, *apm.Span) {
//	span, nextCtx := apm.StartSpan(t.Ctx(), name, spanType)
//	t.ctx = nextCtx
//	t.apmSpans = append(t.apmSpans, span)
//	//if t.apmSpanCtx == nil {
//	//	span, nextCtx := apm.StartSpan(t.Ctx(), name, spanType)
//	//	t.apmSpanCtx = nextCtx
//	//	t.apmSpans = append(t.apmSpans, span)
//	//} else {
//	//	span, nextCtx := apm.StartSpan(t.apmSpanCtx, name, spanType)
//	//	t.apmSpanCtx = nextCtx
//	//	t.apmSpans = append(t.apmSpans, span)
//	//}
//	return t.ctx, span
//}

func (t *Trace) End() {
	for _, s := range t.apmSpans {
		if !s.IsExitSpan() {
			s.End()
		}
	}
	for _, s := range t.spans {
		if s.IsRecording() {
			s.SetStatus(codes.Error, "span not shutdown")
			s.End()
		}
	}
}
