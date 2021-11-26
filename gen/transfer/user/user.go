package user

import (
	"net/http"

	"github.com/fitan/magic/handler/user"
	"github.com/fitan/magic/pkg/types"
)

type SwagBindUserPermissionBody struct {
	RoleID   uint `json:"role_id"`
	DomainID uint `json:"domain_id"`
}

type BindUserPermissionTransfer struct {
}

func (t *BindUserPermissionTransfer) Method() string {
	return http.MethodPost
}

func (t *BindUserPermissionTransfer) Url() string {
	return "/user/:user_id/permission"
}

func (t *BindUserPermissionTransfer) Binder() types.GinXBinder {
	return new(BindUserPermissionBinder)
}

type BindUserPermissionBinder struct {
	val *user.BindUserPermissionIn
}

func (b *BindUserPermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.BindUserPermissionIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagBindUserPermissionBody true " "
// @Param user_id path string true " "
// @Success 200 {object} public.Result{data=string}
// @Router /user/:user_id/permission [post]
func (b *BindUserPermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.BindUserPermission(core, b.val)
}

type SwagUnBindUserPermissionBody struct {
	RoleID   uint `json:"role_id"`
	DomainID uint `json:"domain_id"`
}

type UnBindUserPermissionTransfer struct {
}

func (t *UnBindUserPermissionTransfer) Method() string {
	return http.MethodDelete
}

func (t *UnBindUserPermissionTransfer) Url() string {
	return "/user/:user_id/permission"
}

func (t *UnBindUserPermissionTransfer) Binder() types.GinXBinder {
	return new(UnBindUserPermissionBinder)
}

type UnBindUserPermissionBinder struct {
	val *user.UnBindUserPermissionIn
}

func (b *UnBindUserPermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.UnBindUserPermissionIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagUnBindUserPermissionBody true " "
// @Param user_id path string true " "
// @Success 200 {object} public.Result{data=string}
// @Router /user/:user_id/permission [delete]
func (b *UnBindUserPermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.UnBindUserPermission(core, b.val)
}
