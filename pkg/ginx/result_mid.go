package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"net/http"
	"strconv"
)

type GinXResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

type ResultWrapMid struct {
}

func (r *ResultWrapMid) Forever(core *types.Core) {
	wrapRes := GinXResult{
		Data: core.GinX.BindRes(),
	}
	var code int
	if core.GinX.BindErr() != nil {

		wrapRes.Msg = core.GinX.BindErr().Error()
		code = 5003
	} else {
		code = 2000
	}

	wrapRes.Code = code

	core.Prom.RequestBody(strconv.Itoa(code), core.GinX.GinCtx().Request.Method, core.GinX.GinCtx().FullPath())

	core.GinX.GinCtx().JSON(http.StatusOK, wrapRes)
}

func (r *ResultWrapMid) BindValAfter(core *types.Core) bool {
	return true
	if core.GinX.BindErr() != nil {
		return false
	} else {
		return true
	}
	//if core.GinX.BindErr() != nil {
	//	core.GinX.GinCtx().JSON(http.StatusOK, GinXResult{
	//		Code: 5003,
	//		Msg:  errors.WithMessage(core.GinX.BindErr(), "BindValAfter").Error(),
	//		Data: core.GinX.BindRes(),
	//	})
	//	return false
	//}
	//return true
}

func (r *ResultWrapMid) BindValBefor(core *types.Core) bool {
	return true
	if core.GinX.BindErr() != nil {
		return false
	} else {
		return true
	}
	//if core.GinX.BindErr() != nil {
	//	core.GinX.GinCtx().JSON(http.StatusOK, GinXResult{
	//		Code: 5003,
	//		Msg:  errors.WithMessage(core.GinX.BindErr(), "BindValBefor").Error(),
	//		Data: core.GinX.BindRes(),
	//	})
	//	return false
	//}
	//return true
}

func (r *ResultWrapMid) BindFnAfter(core *types.Core) bool {
	return true
	if core.GinX.BindErr() != nil {
		return false
	} else {
		return true
	}
	//wrapRes := GinXResult{
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
