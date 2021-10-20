package types

type CoreLog interface {
	IsOpenTrace() bool
	TraceLog(spanName string) Logger
}
