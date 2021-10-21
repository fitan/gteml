package core

import (
	"github.com/fitan/gteml/pkg/log"
	"github.com/fitan/gteml/pkg/types"
	"go.uber.org/zap"
)

var xlog *log.Xlog

type CoreLog struct {
	core *types.Context
	xlog *log.Xlog
	//traceLog *log.TraceLog
}

func (c *CoreLog) IsOpenTrace() bool {
	return c.xlog.IsOpenTrace()
}

func (c *CoreLog) TraceLog(spanName string) types.Logger {
	return c.xlog.TraceLog(c.core.Tracer.SpanCtx(spanName))
}

func (c *CoreLog) Log() types.Logger {
	return c.xlog
}

func NewXlog() *log.Xlog {
	if xlog != nil {
		return xlog
	}
	return log.NewXlog(log.WithLogLevel(zap.InfoLevel))
	//return log.NewXlog(log.WithTrace(zap.InfoLevel))
}

type logRegister struct {
}

func (l *logRegister) Reload(c *types.Context) {
	panic("implement me")
}

func (l *logRegister) With(o ...types.Option) types.Register {
	return l
}

func (l *logRegister) Set(c *types.Context) {
	c.CoreLog = &CoreLog{
		core: c,
		xlog: NewXlog(),
	}
}

func (l *logRegister) Unset(c *types.Context) {
	//c.CoreLog = nil
}
