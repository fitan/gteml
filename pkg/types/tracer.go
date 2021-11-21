package types

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Tracer interface {
	SetCtx(ctx context.Context)
	Ctx() context.Context
	SpanCtx(name string) context.Context
	ApmSpanCtx(name string) context.Context
	IsOpen() bool
	End()
	UnSet()
	Tp() *trace.TracerProvider
}
