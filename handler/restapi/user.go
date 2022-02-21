package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/rest"
)

type UserObj struct {
	rest.BaseFieldConf
}

func (u *UserObj) GetTableName() string {
	return "Users"
}

func (u *UserObj) GetModelObj() interface{} {
	return &model.User{}
}

func (u *UserObj) GetModelObjs() interface{} {
	data := make([]model.User, 0, 0)
	return &data
}

func (u *UserObj) GetFirstObj() interface{} {
	return u.GetModelObj()
}

func (u *UserObj) GetFindObj() interface{} {
	return u.GetModelObjs()
}

func (u *UserObj) RelationsField() map[string]rest.RelationFielder {
	return map[string]rest.RelationFielder{"roles": &RolesObj{}, "services": &ServiceObj{}}

}
