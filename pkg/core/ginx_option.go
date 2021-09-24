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

func GinResultWrap(c *Context) {
	res := GinXResult{Data: c.GinX.bindData.res}
	if c.GinX.bindData.err != nil {
		res.Msg = c.GinX.bindData.err.Error()
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

func GinTraceWrap(c *Context) {
	l := c.Log.TraceLog(c.GinX.xkeys[_FnName])
	defer l.End()
	res, _ := json.Marshal(c.GinX.bindData.res)
	val, _ := json.Marshal(c.GinX.bindData.val)
	zf := []zap.Field{zap.String("val", string(val)), zap.String("res", string(res))}
	if c.GinX.bindData.err != nil {
		l.Error(c.GinX.bindData.err.Error(), zf...)
	} else {
		l.Info("handler info", zf...)
	}
}
