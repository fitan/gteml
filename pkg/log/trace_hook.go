package log

import (
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceHook(traceLog *TraceLog) *TraceHook {
	return &TraceHook{
		traceLog: traceLog,
	}
}

type TraceHook struct {
	traceLog *TraceLog
}

func (t TraceHook) Write(p []byte) (n int, err error) {
	t.traceLog.span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("log"), semconv.ExceptionMessageKey.String(string(p))))
	return 0, nil
}
