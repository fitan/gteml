package common

import "context"

type Tracer interface {
	SetCtx(ctx context.Context)
	Ctx() context.Context
	SpanCtx(name string) context.Context
	IsOpen() bool
	End()
}
