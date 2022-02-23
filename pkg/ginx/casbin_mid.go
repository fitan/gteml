package ginx

import (
	"errors"
	"github.com/fitan/magic/pkg/types"
)

type CasbinVerifyer interface {
	ServiceID() (serviceID uint)
}

type CasbinVerifyMid struct {
}

func (c *CasbinVerifyMid) BindValBefore(core *types.Core) bool {
	return true
}

func (c *CasbinVerifyMid) BindValAfter(core *types.Core) bool {
	ginX := core.GetGinX()
	userID, ok := ginX.GinCtx().Get(types.JwtUserIDKey)
	if !ok {
		ginX.SetError(errors.New("userID none of ctx"))
		return false
	}

	userIDuint, ok := userID.(uint)
	if !ok {
		ginX.SetError(errors.New("userID type not is uint"))
		return false
	}

	casbinVerify, ok := ginX.Request().(CasbinVerifyer)
	if !ok {
		ginX.SetError(errors.New("Do not implement CasbinVerifyer"))
		return false
	}

	serviceID := casbinVerify.ServiceID()

	err := core.GetDao().Storage().User().CheckUserPermission(userIDuint, serviceID, ginX.GinCtx().FullPath(), ginX.GinCtx().Request.Method)
	if err != nil {
		ginX.SetError(err)
		return false
	}

	return true
}

func (c *CasbinVerifyMid) BindFnAfter(core *types.Core) bool {
	return true
}

func (c *CasbinVerifyMid) Forever(core *types.Core) {
	return
}
