package log

import (
	"go.elastic.co/apm"
	"go.uber.org/zap"
)

type ApmLog struct {
	*zap.Logger
	*apm.Span
}

func (a *ApmLog) Sync() error {
	a.Span.End()
	return nil
}
