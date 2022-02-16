package restapi

import "github.com/fitan/magic/dao/dal/model"

type RolesObj struct {
}

func (r *RolesObj) GetModelObj() interface{} {
	return &model.Role{}
}

func (r *RolesObj) GetModelObjs() interface{} {
	i := make([]model.Role, 0, 0)
	return i
}

func (r *RolesObj) GetFirstObj() interface{} {
	return r.GetModelObj()
}

func (r *RolesObj) GetFindObj() interface{} {
	return r.GetModelObjs()
}

func NewRolesObj() *RolesObj {
	return &RolesObj{}
}
