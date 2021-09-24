package log

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	//"github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Xlog struct {
	traceLevel zapcore.Level
	tp         *otelsdk.TracerProvider
	*zap.Logger
}

func (x Xlog) TraceLog(ctx context.Context, spanName string) *TraceLog {
	traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	spanCtx, span := x.tp.Tracer(spanName).Start(ctx, spanName)
	traceLog := new(TraceLog)
	hook := NewTraceHook(traceLog)
	traceCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(hook), x.traceLevel)
	wrapCore := zap.WrapCore(
		func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, traceCore)
		})

	traceLog.span = span
	traceLog.ctx = spanCtx
	traceLog.Logger = x.Logger.WithOptions(wrapCore, zap.Fields(zap.String("traceID", traceID)))
	return traceLog
}

type TraceLog struct {
	span trace.Span
	ctx  context.Context
	*zap.Logger
}

func (t *TraceLog) With(fields ...zap.Field) {
	l := t.Logger.With(fields...)
	t.Logger = l
}

func (t *TraceLog) Context() context.Context {
	return t.ctx
}

func (t *TraceLog) End() {
	t.span.End()
}

func (t *TraceLog) Error(msg string, fields ...zap.Field) {
	t.span.SetStatus(codes.Error, "")
	t.Logger.Error(msg, fields...)
}
