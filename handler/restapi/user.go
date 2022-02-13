package restapi

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/rest"
)

type UserObj struct {
}

func (u *UserObj) GetObj() interface{} {
	return &model.User{}
}

func (u *UserObj) GetObjs() interface{} {
	objs := make([]model.User, 0, 0)
	return &objs
}

type User struct {
	*rest.BaseRest
}
