package permission

import (
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
)

type CreatePermissionIn struct {
	Body model.Permission `json:"body"`
}

// @GenApi /permission [post]
func CreatePermission(core *types.Core, in *CreatePermissionIn) (res string, err error) {
	err = core.GetDao().Storage().Permission().Create(&in.Body)
	if err != nil {
		res = "fail"
		return
	}

	res = "succeed"
	return
}

type GetPermissionByIdIn struct {
	Uri struct {
		Id uint `uri:"id"`
	} `json:"uri"`
}

func (g *GetPermissionByIdIn) ServiceID() (serviceID uint) {
	return g.Uri.Id
}

// @GenApi /permission/:id [get]
func GetPermissionById(core *types.Core, in *GetPermissionByIdIn) (res *model.Permission, err error) {
	res, err = core.GetDao().Storage().Permission().GetByID(in.Uri.Id)
	return
}

type DeletePermissionByIdIn struct {
	Uri struct {
		Id uint `uri:"id"`
	} `json:"uri"`
}

// @GenApi /permisssion/:id [delete]
func DeletePermissionById(core *types.Core, in *DeletePermissionByIdIn) (res string, err error) {
	err = core.GetDao().Storage().Permission().DeleteByID(in.Uri.Id)
	if err != nil {
		res = "fail"
		return
	}

	res = "succeed"
	return
}

type UpdatePermissionIn struct {
	Body model.Permission `json:"body"`
}

// @GenApi /permission [put]
func UpdatePermission(core *types.Core, in *UpdatePermissionIn) (res string, err error) {
	err = core.GetDao().Storage().Permission().UpdateById(&in.Body)
	if err != nil {
		res = "fail"
		return
	}

	res = "succeed"
	return
}
