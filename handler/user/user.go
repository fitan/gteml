package user

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/handler/restapi"
	"github.com/fitan/magic/pkg/types"
)

type BindUserPermissionIn struct {
	Uri struct {
		UserID uint `json:"userId" uri:"userId"`
	} `json:"uri"`

	Body struct {
		RoleID   uint `json:"roleId"`
		DomainID uint `json:"domainId"`
	} `json:"body"`
}

// @Description 给用户绑定角色和服务
// @GenApi /user/:userId/permission [post]
func BindUserPermission(core *types.Core, in *BindUserPermissionIn) (string, error) {
	err := core.GetDao().User().BindPermission(in.Uri.UserID, in.Body.RoleID, in.Body.DomainID)
	if err != nil {
		return "fail", err
	}

	return "succeed", nil
}

type UnBindUserPermissionIn struct {
	Uri struct {
		UserID uint `json:"userId" uri:"userId"`
	} `json:"uri"`

	Body struct {
		RoleID   uint `json:"roleId"`
		DomainID uint `json:"domainId"`
	} `json:"body"`
}

// @Description 用户解除绑定
// @GenApi /user/:userId/permission [delete]
func UnBindUserPermission(core *types.Core, in *UnBindUserPermissionIn) (string, error) {
	err := core.GetDao().User().UnBindPermission(in.Uri.UserID, in.Body.RoleID, in.Body.DomainID)
	if err != nil {
		return "fail", err
	}

	return "succeed", nil
}

type GetUserByIDIn struct {
}

// @Description get user by id
// @Tags Users
// @GenApi /users/:id [get]
func GetUserByID(core *types.Core, in *GetUserByIDIn) (*model.User, error) {
	rest := restapi.GetRestfulAll()
	var obj *model.User
	_, err := rest.Users.GetOne(core, obj)
	return obj, err
}

type GetUsersIn struct {
}

type GetUsersOut struct {
	List  *[]model.User `json:"list"`
	Count int64         `json:"count"`
}

// @Description 获取Users
// @Tags Users
// @GenApi /users [get]
func GetUsers(core *types.Core, in *GetUsersIn) (*GetUsersOut, error) {
	rest := restapi.GetRestfulAll()
	var count int64
	objs := make([]model.User, 0)
	_, err := rest.Users.GetList(core, &objs, &count)
	if err != nil {
		return nil, err
	}
	return &GetUsersOut{
		List:  &objs,
		Count: count,
	}, err
}

type UpdateUserIn struct {
}

type UpdateUserOut struct {
}

// @Description UpdateUser
// @Tags Users
// @GenApi /users/:id [put]
func UpdateUser(core *types.Core, in *UpdateUserIn) (*model.User, error) {
	rest := restapi.GetRestfulAll()
	data, err := rest.Users.Update(core)
	if err != nil {
		return nil, err
	}
	return data.(*model.User), nil
}

type DeleteUserIn struct {
}

type DeleteUserOut struct {
}

// @Description DeleteUser
// @Tags Users
// @GenApi /users/:id [delete]
func DeleteUser(core *types.Core, in *DeleteUserIn) (*model.User, error) {
	rest := restapi.GetRestfulAll()
	data, err := rest.Users.Delete(core)
	if err != nil {
		return nil, err
	}
	return data.(*model.User), nil
}

type CreateUserIn struct {
}

type CreateUserOut struct {
}

// @Description CreateUser
// @Tags Users
// @GenApi /users [post]
func CreateUser(core *types.Core, in *CreateUserIn) (*model.User, error) {
	data, err := restapi.GetRestfulAll().Users.Create(core)
	if err != nil {
		return nil, err
	}
	return data.(*model.User), nil
}
