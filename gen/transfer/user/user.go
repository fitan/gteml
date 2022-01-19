package user

import (
	"net/http"

	"github.com/fitan/magic/handler/user"
	"github.com/fitan/magic/pkg/types"
)

type SwagBindUserPermissionBody struct {
	RoleID   uint `json:"roleId"`
	DomainID uint `json:"domainId"`
}

type BindUserPermissionTransfer struct {
}

func (t *BindUserPermissionTransfer) Method() string {
	return http.MethodPost
}

func (t *BindUserPermissionTransfer) Url() string {
	return "/user/:userId/permission"
}

func (t *BindUserPermissionTransfer) FuncName() string {
	return "user.BindUserPermission"
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
// @Param userId path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description 给用户绑定角色和服务
// @Router /user/:userId/permission [post]
func (b *BindUserPermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.BindUserPermission(core, b.val)
}

type SwagUnBindUserPermissionBody struct {
	RoleID   uint `json:"roleId"`
	DomainID uint `json:"domainId"`
}

type UnBindUserPermissionTransfer struct {
}

func (t *UnBindUserPermissionTransfer) Method() string {
	return http.MethodDelete
}

func (t *UnBindUserPermissionTransfer) Url() string {
	return "/user/:userId/permission"
}

func (t *UnBindUserPermissionTransfer) FuncName() string {
	return "user.UnBindUserPermission"
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
// @Param userId path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description 用户解除绑定
// @Router /user/:userId/permission [delete]
func (b *UnBindUserPermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.UnBindUserPermission(core, b.val)
}

type GetUserByIDTransfer struct {
}

func (t *GetUserByIDTransfer) Method() string {
	return http.MethodGet
}

func (t *GetUserByIDTransfer) Url() string {
	return "/user/:id"
}

func (t *GetUserByIDTransfer) FuncName() string {
	return "user.GetUserByID"
}

func (t *GetUserByIDTransfer) Binder() types.GinXBinder {
	return new(GetUserByIDBinder)
}

type GetUserByIDBinder struct {
	val *user.GetUserByIDIn
}

func (b *GetUserByIDBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.GetUserByIDIn{}
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
// @Success 200 {object} ginx.GinXResult{data=model.User}
// @Description get user by id
// @Router /user/:id [get]
func (b *GetUserByIDBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.GetUserByID(core, b.val)
}
