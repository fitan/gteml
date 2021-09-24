package core

import (
	"github.com/fitan/gteml/pkg/log"
	"github.com/fitan/gteml/pkg/trace"
	"go.uber.org/zap"
)

var xlog *log.Xlog

type CoreLog struct {
	core *Context
	xlog *log.Xlog
}

func (c *CoreLog) TraceLog(spanName string) *log.TraceLog {
	if c.core.TraceLog == nil {
		c.core.TraceLog = c.xlog.TraceLog(c.core.Ctx, spanName)
		return c.core.TraceLog
	} else {
		c.core.TraceLog = c.xlog.TraceLog(c.core.TraceLog.Context(), spanName)
		return c.core.TraceLog
	}
}

func NewXlog() *log.Xlog {
	if xlog != nil {
		return xlog
	}
	return log.NewXlog(log.WithTrace(trace.GetTp(), zap.InfoLevel))
}

type logRegister struct {
}

func (l logRegister) Set(c *Context) {
	c.Log = &CoreLog{
		core: c,
		xlog: NewXlog(),
	}
}

func (l logRegister) Unset(c *Context) {
	c.Log = nil
}
