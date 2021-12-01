package types

import (
	"context"
	"go.elastic.co/apm"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Tracer interface {
	SetCtx(ctx context.Context)
	Ctx() context.Context
	SpanCtx(name string) context.Context
	ApmSpanCtx(name string, spanType string) (context.Context, *apm.Span)
	IsOpen() bool
	End()
	UnSet()
	Tp() *trace.TracerProvider
}
