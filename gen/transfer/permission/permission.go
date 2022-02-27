package permission

import (
	"net/http"

	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/handler/permission"
	"github.com/fitan/magic/pkg/types"
)

type SwagCreatePermissionBody model.Permission

type CreatePermissionTransfer struct {
}

func (t *CreatePermissionTransfer) Method() string {
	return http.MethodPost
}

func (t *CreatePermissionTransfer) Url() string {
	return "/permission"
}

func (t *CreatePermissionTransfer) FuncName() string {
	return "permission.CreatePermission"
}

func (t *CreatePermissionTransfer) Binder() types.GinXBinder {
	return new(CreatePermissionBinder)
}

type CreatePermissionBinder struct {
	val *permission.CreatePermissionIn
}

func (b *CreatePermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &permission.CreatePermissionIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagCreatePermissionBody true " "
// @Success 200 {object} ginx.XResult{data=string}
// @Router /permission [post]
func (b *CreatePermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return permission.CreatePermission(core, b.val)
}

type GetPermissionByIdTransfer struct {
}

func (t *GetPermissionByIdTransfer) Method() string {
	return http.MethodGet
}

func (t *GetPermissionByIdTransfer) Url() string {
	return "/permission/:id"
}

func (t *GetPermissionByIdTransfer) FuncName() string {
	return "permission.GetPermissionById"
}

func (t *GetPermissionByIdTransfer) Binder() types.GinXBinder {
	return new(GetPermissionByIdBinder)
}

type GetPermissionByIdBinder struct {
	val *permission.GetPermissionByIdIn
}

func (b *GetPermissionByIdBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &permission.GetPermissionByIdIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} ginx.XResult{data=model.Permission}
// @Router /permission/{id} [get]
func (b *GetPermissionByIdBinder) BindFn(core *types.Core) (interface{}, error) {
	return permission.GetPermissionById(core, b.val)
}

type DeletePermissionByIdTransfer struct {
}

func (t *DeletePermissionByIdTransfer) Method() string {
	return http.MethodDelete
}

func (t *DeletePermissionByIdTransfer) Url() string {
	return "/permisssion/:id"
}

func (t *DeletePermissionByIdTransfer) FuncName() string {
	return "permission.DeletePermissionById"
}

func (t *DeletePermissionByIdTransfer) Binder() types.GinXBinder {
	return new(DeletePermissionByIdBinder)
}

type DeletePermissionByIdBinder struct {
	val *permission.DeletePermissionByIdIn
}

func (b *DeletePermissionByIdBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &permission.DeletePermissionByIdIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} ginx.XResult{data=string}
// @Router /permisssion/{id} [delete]
func (b *DeletePermissionByIdBinder) BindFn(core *types.Core) (interface{}, error) {
	return permission.DeletePermissionById(core, b.val)
}

type SwagUpdatePermissionBody model.Permission

type UpdatePermissionTransfer struct {
}

func (t *UpdatePermissionTransfer) Method() string {
	return http.MethodPut
}

func (t *UpdatePermissionTransfer) Url() string {
	return "/permission"
}

func (t *UpdatePermissionTransfer) FuncName() string {
	return "permission.UpdatePermission"
}

func (t *UpdatePermissionTransfer) Binder() types.GinXBinder {
	return new(UpdatePermissionBinder)
}

type UpdatePermissionBinder struct {
	val *permission.UpdatePermissionIn
}

func (b *UpdatePermissionBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &permission.UpdatePermissionIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagUpdatePermissionBody true " "
// @Success 200 {object} ginx.XResult{data=string}
// @Router /permission [put]
func (b *UpdatePermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return permission.UpdatePermission(core, b.val)
}
