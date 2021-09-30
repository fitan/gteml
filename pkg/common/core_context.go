package common

import "github.com/fitan/gteml/pkg/log"

type Context interface {
	Tracer
	TraceLog(spanName string) *log.TraceLog
}
