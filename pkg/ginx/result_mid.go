package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strconv"
)

var SkipWrapError = errors.New("Skip ResultWrapMid")

type XResult struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
	Data    interface{} `json:"data"`
}

type ResultWrapMid struct {
}

func (r *ResultWrapMid) Forever(core *types.Core) {
	//err := recover()
	//if err != nil {
	//	const size = 64 << 10
	//	buf := make([]byte, size)
	//	buf = buf[:runtime.Stack(buf, false)]
	//	log := core.GetCoreLog().ApmLog("pkg.ginx.wrapMid.recover")
	//	log.Error("panic", zap.Any("err", err), zap.String("stack", string(buf)))
	//	log.Sync()
	//
	//	core.GinX.SetError(errors.New("系统错误，请联系管理员"))
	//}

	if errors.Is(core.GinX.LastError(), SkipWrapError) {
		return
	}

	trace.SpanFromContext(core.GetTrace().Ctx()).SpanContext().TraceID()
	wrapRes := XResult{
		Data:    core.GinX.Response(),
		TraceId: trace.SpanFromContext(core.GetTrace().Ctx()).SpanContext().TraceID().String(),
	}
	var code int
	if core.GinX.LastError() != nil {

		wrapRes.Msg = core.GinX.LastError().Error()
		code = 5003
	} else {
		code = 2000
	}

	wrapRes.Code = code

	core.Prom.RequestBody(strconv.Itoa(code), core.GinX.GinCtx().Request.Method, core.GinX.GinCtx().FullPath())

	core.GinX.GinCtx().JSON(http.StatusOK, wrapRes)
}

func (r *ResultWrapMid) BindValAfter(core *types.Core) bool {
	if core.GinX.LastError() != nil {
		return false
	} else {
		return true
	}
	//if core.GinX.BindErr() != nil {
	//	core.GinX.GinCtx().JSON(http.StatusOK, XResult{
	//		Code: 5003,
	//		Msg:  errors.WithMessage(core.GinX.BindErr(), "BindValAfter").Error(),
	//		Data: core.GinX.BindRes(),
	//	})
	//	return false
	//}
	//return true
}

func (r *ResultWrapMid) BindValBefore(core *types.Core) bool {
	if core.GinX.LastError() != nil {
		return false
	} else {
		return true
	}
	//if core.GinX.BindErr() != nil {
	//	core.GinX.GinCtx().JSON(http.StatusOK, XResult{
	//		Code: 5003,
	//		Msg:  errors.WithMessage(core.GinX.BindErr(), "BindValBefor").Error(),
	//		Data: core.GinX.BindRes(),
	//	})
	//	return false
	//}
	//return true
}

func (r *ResultWrapMid) BindFnAfter(core *types.Core) bool {
	if core.GinX.LastError() != nil {
		return false
	} else {
		return true
	}
	//wrapRes := XResult{
	//	Msg: errors.WithMessage(core.GinX.BindErr(), "BindFnAfter").Error(),
	//	Data: core.GinX.BindRes(),
	//}
	//if core.GinX.BindErr() != nil {
	//	wrapRes.Code = 5003
	//} else {
	//	wrapRes.Code = 2000
	//}
	//
	//core.GinX.GinCtx().JSON(http.StatusOK, wrapRes)
	//return true
}
