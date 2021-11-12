package ginx

import (
	"github.com/fitan/magic/pkg/types"
)

const FnName = "fn_name"

func WithEntryMid(m *[]types.Middleware) GinXOption {
	return func(g *GinX) {
		g.SetEntryMid(m)
	}
}

func WithHandlerMid(mids ...types.Middleware) GinXHandlerOption {
	ms := make([]types.Middleware, 0, len(mids))
	ms = append(ms, mids...)
	return func(core *types.Core) error {
		core.GinX.SetHandlerMid(&ms)
		return nil
	}
}

//
//func GinXResultWrap(c *types.Core)  {
//	res := GinXResult{Data: c.GinX.BindRes()}
//	if c.GinX.BindErr() != nil {
//		res.Msg = c.GinX.BindErr().Error()
//		res.Code = 5003
//		c.GinX.GinCtx().JSON(http.StatusInternalServerError, res)
//	}
//
//	res.Code = 2000
//	c.GinX.GinCtx().JSON(http.StatusOK, res)
//}
//
//type GinXResult struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg,omitempty"`
//	Data interface{} `json:"data"`
//}
//
//func GinXTraceWrap(c *types.Core) {
//	if !c.Tracer.IsOpen() {
//		return
//	}
//	l := c.CoreLog.TraceLog("GinXTraceWrap")
//	defer l.Sync()
//	res, _ := json.Marshal(c.GinX.BindRes())
//	req, _ := json.Marshal(c.GinX.BindReq())
//	zf := []zap.Field{zap.String("req", string(req)), zap.String("res", string(res))}
//	if c.GinX.BindErr() != nil {
//		l.Error(c.GinX.BindErr().Error(), zf...)
//	} else {
//		l.Info("handler info", zf...)
//	}
//}
