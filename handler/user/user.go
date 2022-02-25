package user

import (
	"github.com/fitan/magic/dao/dal/model"
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
	Uri struct {
		ID uint `uri:"id"`
	}
}

// @Description get user by id
// @GenApi /user/:id [get]
func GetUserByID(core *types.Core, in *GetUserByIDIn) (*model.User, error) {
	return core.GetDao().User().ById(in.Uri.ID)
}
