package context

import (
	"github.com/fitan/magic/pkg/log"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap/zapcore"
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

type logRegister struct {
	xlog *log.Xlog
}

func (l *logRegister) GetXlog() *log.Xlog {
	if l.xlog == nil {
		l.xlog = log.NewXlog(log.WithLogLevel(zapcore.Level(Conf.MyConf.Log.Lervel)), log.WithTrace(zapcore.Level(Conf.MyConf.Log.TraceLervel)), log.WithLogFileName(Conf.MyConf.Log.FileName, Conf.MyConf.Log.Dir))
	}
	return l.xlog
}

func (l *logRegister) Reload(c *types.Context) {
	l.xlog = nil
}

func (l *logRegister) With(o ...types.Option) types.Register {
	return l
}

func (l *logRegister) Set(c *types.Context) {
	c.CoreLog = &CoreLog{
		core: c,
		xlog: l.GetXlog(),
	}
}

func (l *logRegister) Unset(c *types.Context) {
	//c.CoreLog = nil
}
