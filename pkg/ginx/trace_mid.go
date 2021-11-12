package ginx

import (
	"encoding/json"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

type TraceMid struct {
}

func (t *TraceMid) Forever(core *types.Core) {
	if !core.Tracer.IsOpen() {
		return
	}
	l := core.CoreLog.TraceLog("GinXTraceWrap")
	defer l.Sync()
	res, _ := json.Marshal(core.GinX.BindRes())
	req, _ := json.Marshal(core.GinX.BindReq())
	zf := []zap.Field{zap.String("req", string(req)), zap.String("res", string(res))}
	if core.GinX.BindErr() != nil {
		l.Error(core.GinX.BindErr().Error(), zf...)
	} else {
		l.Info("handler info", zf...)
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
