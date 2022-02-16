package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/rest"
)

type UserObj struct {
}

func (u *UserObj) GetTableName() string {
	return "users"
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

func (u *UserObj) GetUpdateObj() interface{} {
	return u.GetModelObj()
}

func (u *UserObj) GetCreateObj() interface{} {
	return u.GetModelObj()
}

type User struct {
	*rest.BaseRest
}
