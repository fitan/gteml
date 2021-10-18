package core

import (
	"github.com/fitan/gteml/pkg/log"
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
	return c.xlog.TraceLog(c.core.SpanCtx(spanName))
}

func NewXlog() *log.Xlog {
	if xlog != nil {
		return xlog
	}
	return log.NewXlog()
	//return log.NewXlog(log.WithTrace(zap.InfoLevel))
}

type logRegister struct {
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
