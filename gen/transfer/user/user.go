package user

import (
	"github.com/fitan/magic/dao/dal/model"
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
// @Success 200 {object} ginx.XResult{data=string}
// @Description 给用户绑定角色和服务
// @Router /user/{userId}/permission [post]
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
// @Success 200 {object} ginx.XResult{data=string}
// @Description 用户解除绑定
// @Router /user/{userId}/permission [delete]
func (b *UnBindUserPermissionBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.UnBindUserPermission(core, b.val)
}

type GetUserByIDTransfer struct {
}

func (t *GetUserByIDTransfer) Method() string {
	return http.MethodGet
}

func (t *GetUserByIDTransfer) Url() string {
	return "/users/:id"
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

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=model.User}
// @Description get user by id
// @Tags Users
// @Router /users/{id} [get]
func (b *GetUserByIDBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.GetUserByID(core, b.val)
}

type GetUsersTransfer struct {
}

func (t *GetUsersTransfer) Method() string {
	return http.MethodGet
}

func (t *GetUsersTransfer) Url() string {
	return "/users"
}

func (t *GetUsersTransfer) FuncName() string {
	return "user.GetUsers"
}

func (t *GetUsersTransfer) Binder() types.GinXBinder {
	return new(GetUsersBinder)
}

type GetUsersBinder struct {
	val *user.GetUsersIn
}

func (b *GetUsersBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.GetUsersIn{}
	b.val = in

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=user.GetUsersOut}
// @Description 获取Users
// @Tags Users
// @Router /users [get]
func (b *GetUsersBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.GetUsers(core, b.val)
}

type UpdateUserTransfer struct {
}

func (t *UpdateUserTransfer) Method() string {
	return http.MethodPut
}

func (t *UpdateUserTransfer) Url() string {
	return "/users/:id"
}

func (t *UpdateUserTransfer) FuncName() string {
	return "user.UpdateUser"
}

func (t *UpdateUserTransfer) Binder() types.GinXBinder {
	return new(UpdateUserBinder)
}

type UpdateUserBinder struct {
	val *user.UpdateUserIn
}

func (b *UpdateUserBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.UpdateUserIn{}
	b.val = in

	return b.val, err
}

var _ = model.User{}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=model.User}
// @Description UpdateUser
// @Tags Users
// @Router /users/{id} [put]
func (b *UpdateUserBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.UpdateUser(core, b.val)
}

type DeleteUserTransfer struct {
}

func (t *DeleteUserTransfer) Method() string {
	return http.MethodDelete
}

func (t *DeleteUserTransfer) Url() string {
	return "/users/:id"
}

func (t *DeleteUserTransfer) FuncName() string {
	return "user.DeleteUser"
}

func (t *DeleteUserTransfer) Binder() types.GinXBinder {
	return new(DeleteUserBinder)
}

type DeleteUserBinder struct {
	val *user.DeleteUserIn
}

func (b *DeleteUserBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.DeleteUserIn{}
	b.val = in

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=model.User}
// @Description DeleteUser
// @Tags Users
// @Router /users/{id} [delete]
func (b *DeleteUserBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.DeleteUser(core, b.val)
}

type CreateUserTransfer struct {
}

func (t *CreateUserTransfer) Method() string {
	return http.MethodPost
}

func (t *CreateUserTransfer) Url() string {
	return "/users"
}

func (t *CreateUserTransfer) FuncName() string {
	return "user.CreateUser"
}

func (t *CreateUserTransfer) Binder() types.GinXBinder {
	return new(CreateUserBinder)
}

type CreateUserBinder struct {
	val *user.CreateUserIn
}

func (b *CreateUserBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.CreateUserIn{}
	b.val = in

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=model.User}
// @Description CreateUser
// @Tags Users
// @Router /users [post]
func (b *CreateUserBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.CreateUser(core, b.val)
}
