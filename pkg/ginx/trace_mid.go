package ginx

import (
	"encoding/json"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

type TraceMid struct {
}

func (t *TraceMid) Forever(core *types.Core) {

	send := func() {
		l := core.CoreLog.TraceLog("GinXTraceWrap")
		defer l.Sync()
		res, _ := json.Marshal(core.GinX.Response())
		req, _ := json.Marshal(core.GinX.Request())
		zf := []zap.Field{zap.String("req", string(req)), zap.String("res", string(res))}
		if core.GinX.LastError() != nil {
			l.Error(core.GinX.LastError().Error(), zf...)
		} else {
			l.Info("handler info", zf...)
		}
	}

	if core.Tracer.IsOpen() {
		if core.Config.GetMyConf().Log.TraceLervel < 2 {
			send()
			return
		}

		if core.GinX.LastError() != nil {
			send()
			return
		}

	}
}

func (t *TraceMid) BindValAfter(core *types.Core) bool {
	return true
}

func (t *TraceMid) BindValBefor(core *types.Core) bool {
	return true
}

func (t *TraceMid) BindFnAfter(core *types.Core) bool {
	return true
}
