package core

import (
	"github.com/fitan/magic/pkg/log"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap/zapcore"
)

var xlog *log.Xlog

type CoreLog struct {
	core     *types.Core
	xlog     *log.Xlog
	traceLog *log.TraceLog
}

func (c *CoreLog) IsOpenTrace() bool {
	return c.xlog.IsOpenTrace()
}

func (c *CoreLog) TraceLog(spanName string) types.Logger {
	return c.xlog.TraceLog(c.core.Tracer.SpanCtx(spanName))
	//ctx, _ := c.core.Tracer.ApmSpanCtx(spanName, "method")
	//return c.xlog.ApmLog(ctx)
}

//func (c *CoreLog) TraceLog(spanName string) types.Logger {
//	return c.xlog.TraceLog(c.core.Tracer.SpanCtx(spanName))
//}

func (c *CoreLog) Log() types.Logger {
	return c.xlog
}

type LogRegister struct {
	xlog *log.Xlog
}

func (l *LogRegister) GetXlog() *log.Xlog {
	if l.xlog == nil {
		l.xlog = log.NewXlog(
			log.WithLogLevel(zapcore.Level(ConfReg.Confer.GetMyConf().Log.Lervel)),
			log.WithTrace(zapcore.Level(ConfReg.Confer.GetMyConf().Log.TraceLervel), map[string]struct{}{"carry": {}}),
			log.WithLogFileName(ConfReg.Confer.GetMyConf().Log.FileName, ConfReg.Confer.GetMyConf().Log.Dir))
	}
	return l.xlog
}

func (l *LogRegister) Reload(c *types.Core) {
	l.xlog = nil
}

func (l *LogRegister) With(o ...types.Option) types.Register {
	return l
}

func (l *LogRegister) Set(c *types.Core) {
	c.CoreLog = &CoreLog{
		core: c,
		xlog: l.GetXlog(),
	}
}

func (l *LogRegister) Unset(c *types.Core) {
	//c.CoreLog = nil
}
