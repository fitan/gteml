package core

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

const _FnName = "fn_name"

func WithWrap(ops ...Option) GinOption {
	wrap := make([]Option, 0, len(ops))
	for _, o := range ops {
		wrap = append(wrap, o)
	}
	return func(g *GinX) {
		g.resultWrap = wrap
	}
}

func GinXResultWrap(c *Context) {
	res := GinXResult{Data: c.GinX.bindRes}
	if c.GinX.bindErr != nil {
		res.Msg = c.GinX.bindErr.Error()
		res.Code = 5003
		c.GinX.Context.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = 2000
	c.GinX.Context.JSON(http.StatusOK, res)
}

type GinXResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

func GinXTraceWrap(c *Context) {
	if !c.Tracer.IsOpen() {
		return
	}
	l := c.TraceLog("GinxTraceWrap")
	defer l.Sync()
	res, _ := json.Marshal(c.GinX.bindRes)
	val, _ := json.Marshal(c.GinX.bindVal)
	zf := []zap.Field{zap.String("val", string(val)), zap.String("res", string(res))}
	if c.GinX.bindErr != nil {
		l.Error(c.GinX.bindErr.Error(), zf...)
	} else {
		l.Info("handler info", zf...)
	}
}
