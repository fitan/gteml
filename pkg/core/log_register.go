package core

import (
	"github.com/fitan/gteml/pkg/log"
	"go.uber.org/zap"
)

var xlog *log.Xlog

type CoreLog struct {
	core *Context
	xlog *log.Xlog
	//traceLog *log.TraceLog
}

func (c *CoreLog) IsOpenTrace() bool {
	return c.xlog.IsOpenTrace()
}

func (c *CoreLog) TraceLog(spanName string) *log.TraceLog {
	return c.xlog.TraceLog(c.core.Tracer.SpanCtx(spanName))
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

func (l *logRegister) Reload(c *Context) {
	panic("implement me")
}

func (l *logRegister) With(o ...Option) Register {
	return l
}

func (l *logRegister) Set(c *Context) {
	c.CoreLog = &CoreLog{
		core: c,
		xlog: NewXlog(),
	}
}

func (l *logRegister) Unset(c *Context) {
	//c.CoreLog = nil
}
