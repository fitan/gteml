package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/rest"
)

type RolesObj struct {
}

func (r *RolesObj) GetTableName() string {
	return "Roles"
}

func (r *RolesObj) GetModelObj() interface{} {
	return &model.Role{}
}

func (r *RolesObj) GetModelObjs() interface{} {
	i := make([]model.Role, 0, 0)
	return &i
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

type RolesFieldConf struct {
	rest.BaseFieldConf
}

func (r *RolesFieldConf) RelationField() map[string]rest.RelationFielder {
	return nil
}
