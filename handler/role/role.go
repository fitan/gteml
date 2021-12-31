package role

import (
	"github.com/fitan/magic/pkg/types"
)

type RolePermissionIn struct {
	Uri struct {
		RoleID uint `json:"roleId" uri:"roleId"`
	} `json:"uri"`

	Body struct {
		PermissionID uint `json:"permissionId"`
	} `json:"body"`
}

// @GenApi /role/:roleId/permission [post]
func BindRolePermission(core *types.Core, in *RolePermissionIn) (string, error) {
	err := core.GetDao().Storage().Role().BindPermission(in.Uri.RoleID, in.Body.PermissionID)
	if err != nil {
		return "fail", err
	}

	return "succeed", err
}

// @GenApi /role/:roleId/permission [delete]
func UnBindRolePermission(core *types.Core, in *RolePermissionIn) (string, error) {
	err := core.GetDao().Storage().Role().UnBindPermission(in.Uri.RoleID, in.Body.PermissionID)
	if err != nil {
		return "fail", err
	}

	return "succeed", err
}
