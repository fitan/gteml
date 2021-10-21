package types

type CoreLoger interface {
	IsOpenTrace() bool
	TraceLog(spanName string) Logger
	Log() Logger
}
