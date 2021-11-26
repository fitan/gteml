package role

import (
	"github.com/fitan/magic/pkg/types"
)

type RolePermissionIn struct {
	Uri struct {
		RoleID uint `uri:"role_id"`
	} `json:"uri"`

	Body struct {
		PermissionID uint `json:"permission_id"`
	} `json:"body"`
}

// @GenApi /role/:role_id/permission [post]
func BindRolePermission(core *types.Core, in *RolePermissionIn) (string, error) {
	err := core.GetDao().Storage().Role().BindPermission(in.Uri.RoleID, in.Body.PermissionID)
	if err != nil {
		return "fail", err
	}

	return "succeed", err
}

// @GenApi /role/:role_id/permission [delete]
func UnBindRolePermission(core *types.Core, in *RolePermissionIn) (string, error) {
	err := core.GetDao().Storage().Role().UnBindPermission(in.Uri.RoleID, in.Body.PermissionID)
	if err != nil {
		return "fail", err
	}

	return "succeed", err
}
