package ginx

import (
	"encoding/json"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
	"net/http"
)

const FnName = "fn_name"

func WithWrap(ops ...types.Option) GinOption {
	wrap := make([]types.Option, 0, len(ops))
	for _, o := range ops {
		wrap = append(wrap, o)
	}
	return func(g *GinX) {
		g.resultWrap = wrap
	}
}

func GinXResultWrap(c *types.Context) {
	res := GinXResult{Data: c.GinX.BindRes()}
	if c.GinX.BindErr() != nil {
		res.Msg = c.GinX.BindErr().Error()
		res.Code = 5003
		c.GinX.GinCtx().JSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = 2000
	c.GinX.GinCtx().JSON(http.StatusOK, res)
}

type GinXResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

func GinXTraceWrap(c *types.Context) {
	if !c.Tracer.IsOpen() {
		return
	}
	l := c.CoreLog.TraceLog("GinXTraceWrap")
	defer l.Sync()
	res, _ := json.Marshal(c.GinX.BindRes())
	req, _ := json.Marshal(c.GinX.BindReq())
	zf := []zap.Field{zap.String("req", string(req)), zap.String("res", string(res))}
	if c.GinX.BindErr() != nil {
		l.Error(c.GinX.BindErr().Error(), zf...)
	} else {
		l.Info("handler info", zf...)
	}
}
