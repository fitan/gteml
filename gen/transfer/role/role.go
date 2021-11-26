package role

import (
	"net/http"

	"github.com/fitan/magic/handler/role"
	"github.com/fitan/magic/pkg/types"
)

type SwagBindRolePermissionBody struct {
	PermissionID uint `json:"permission_id"`
}

type BindRolePermissionTransfer struct {
}

func (t *BindRolePermissionTransfer) Method() string {
	return http.MethodPost
}

func (t *BindRolePermissionTransfer) Url() string {
	return "/role/:role_id/permission"
}

func (t *BindRolePermissionTransfer) Binder() types.GinXBinder {
	return new(BindRolePermissionBinder)
}

type BindRolePermissionBinder struct {
	val *role.RolePermissionIn
}

func (b *BindRolePermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &role.RolePermissionIn{}
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
// @Param body body SwagBindRolePermissionBody true " "
// @Param role_id path string true " "
// @Success 200 {object} public.Result{data=string}
// @Router /role/:role_id/permission [post]
func (b *BindRolePermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return role.BindRolePermission(core, b.val)
}

type SwagUnBindRolePermissionBody struct {
	PermissionID uint `json:"permission_id"`
}

type UnBindRolePermissionTransfer struct {
}

func (t *UnBindRolePermissionTransfer) Method() string {
	return http.MethodDelete
}

func (t *UnBindRolePermissionTransfer) Url() string {
	return "/role/:role_id/permission"
}

func (t *UnBindRolePermissionTransfer) Binder() types.GinXBinder {
	return new(UnBindRolePermissionBinder)
}

type UnBindRolePermissionBinder struct {
	val *role.RolePermissionIn
}

func (b *UnBindRolePermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &role.RolePermissionIn{}
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
// @Param body body SwagUnBindRolePermissionBody true " "
// @Param role_id path string true " "
// @Success 200 {object} public.Result{data=string}
// @Router /role/:role_id/permission [delete]
func (b *UnBindRolePermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return role.UnBindRolePermission(core, b.val)
}
