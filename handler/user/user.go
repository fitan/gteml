package user

import "github.com/fitan/magic/pkg/types"

type BindUserPermissionIn struct {
	Uri struct {
		UserID uint `uri:"user_id"`
	} `json:"uri"`

	Body struct {
		RoleID   uint `json:"role_id"`
		DomainID uint `json:"domain_id"`
	} `json:"body"`
}

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

// @GenApi /user/:user_id/permission [delete]
func UnBindUserPermission(core *types.Core, in *UnBindUserPermissionIn) (string, error) {
	err := core.GetDao().Storage().User().UnBindPermission(in.Uri.UserID, in.Body.RoleID, in.Body.DomainID)
	if err != nil {
		return "fail", err
	}

	return "succeed", nil
}
