package user

import (
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
)

type BindUserPermissionIn struct {
	Uri struct {
		UserID uint `uri:"user_id"`
	} `json:"uri"`

	Body struct {
		RoleID   uint `json:"role_id"`
		DomainID uint `json:"domain_id"`
	} `json:"body"`
}

// @Description 给用户绑定角色和服务
// @GenApi /user/:user_id/permission [post]
func BindUserPermission(core *types.Core, in *BindUserPermissionIn) (string, error) {
	err := core.GetDao().Storage().User().BindPermission(in.Uri.UserID, in.Body.RoleID, in.Body.DomainID)
	if err != nil {
		return "fail", err
	}

	return "succeed", nil
}

type UnBindUserPermissionIn struct {
	Uri struct {
		UserID uint `uri:"user_id"`
	} `json:"uri"`

	Body struct {
		RoleID   uint `json:"role_id"`
		DomainID uint `json:"domain_id"`
	}
}

// @Description 用户解除绑定
// @GenApi /user/:user_id/permission [delete]
func UnBindUserPermission(core *types.Core, in *UnBindUserPermissionIn) (string, error) {
	err := core.GetDao().Storage().User().UnBindPermission(in.Uri.UserID, in.Body.RoleID, in.Body.DomainID)
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
	return core.GetDao().Storage().User().ById(in.Uri.ID)
}
