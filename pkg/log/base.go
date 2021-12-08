package log

import (
	"context"
	"github.com/fitan/magic/pkg/apm/apmzap"
	"go.elastic.co/apm"
	"go.opentelemetry.io/otel/codes"
	//otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	//"github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Xlog struct {
	traceLevel zapcore.Level
	openTrace  bool
	//tp         *otelsdk.TracerProvider
	*zap.Logger
	filter map[string]struct{}
}

func (x *Xlog) IsOpenTrace() bool {
	return x.openTrace
}

func (x *Xlog) ApmLog(ctx context.Context) *ApmLog {
	span := apm.SpanFromContext(ctx)
	return &ApmLog{
		Logger: x.Logger.WithOptions(zap.WrapCore((&apmzap.Core{Filter: x.filter}).WrapCore)).With(apmzap.TraceContext(ctx)...),
		Span:   span,
	}
}

func (x *Xlog) TraceLog(ctx context.Context) *TraceLog {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()
	//traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	//spanCtx, span := x.tp.Tracer(spanName).Start(ctx, spanName)
	traceLog := new(TraceLog)
	hook := NewTraceHook(traceLog)
	traceCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(hook), x.traceLevel)
	wrapCore := zap.WrapCore(
		func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, traceCore)
		})

	//traceLog.root = x
	traceLog.span = span
	traceLog.ctx = ctx
	traceLog.Logger = x.Logger.WithOptions(wrapCore, zap.Fields(zap.String("traceID", traceID)))
	return traceLog
}

type TraceLog struct {
	span trace.Span
	ctx  context.Context
	//root *Xlog
	*zap.Logger
}

//func (t *TraceLog) NextTraceLog(spanName string) *TraceLog {
//	traceID := t.span.SpanContext().TraceID().String()
//	spanCtx, span := t.root.tp.Tracer(spanName).Start(t.ctx, spanName)
//	traceLog := new(TraceLog)
//	hook := NewTraceHook(traceLog)
//	traceCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(hook), t.root.traceLevel)
//	wrapCore := zap.WrapCore(
//		func(core zapcore.Core) zapcore.Core {
//			return zapcore.NewTee(core, traceCore)
//		})
//
//	traceLog.root = t.root
//	traceLog.span = span
//	traceLog.ctx = spanCtx
//	traceLog.Logger = t.root.Logger.WithOptions(wrapCore, zap.Fields(zap.String("traceID", traceID)))
//	return traceLog
//}

//func (t *TraceLog) Context() context.Context {
//	return t.ctx
//}

//Sync 为了统一interface span.end() 别名
func (t *TraceLog) Sync() error {
	t.span.End()
	return nil
	//t.Logger.Sync()
}

//Error 增加trace的error状态
func (t *TraceLog) Error(msg string, fields ...zap.Field) {
	t.span.SetStatus(codes.Error, msg)
	t.Logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}
