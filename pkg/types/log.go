package types

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CoreLoger interface {
	IsOpenTrace() bool
	TraceLog(spanName string) Logger
	Log() Logger
	ApmLog(spanName string) *zap.Logger
	//ApmLog(spanName string) Logger
}

type Logger interface {
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
	Core() zapcore.Core
}
