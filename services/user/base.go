package user

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

func NewUser(core types.ServiceCore) *User {
	return &User{core}
}

type User struct {
	Core types.ServiceCore
}

func (u *User) Login(username string, password string) (*model.User, error) {
	return u.Core.GetDao().User().CheckPassword(username, password)
}

func (u *User) ModifyPassword(id int, password string) error {
	return u.Core.GetDao().Native().User.ModifyPassword(id, password)
}

func (u *User) Read() string {
	//log := u.Core.GetCoreLog().TraceLog("user.read")
	//defer log.Sync()
	//
	//log := u.Core.GetCoreLog().ApmLog("user.read")
	req, _ := u.Core.GetDao().User().ById(1)

	log := u.Core.GetCoreLog().TraceLog("read")
	log.Error("this is read", zap.String("read", "read"), zap.Any("carry", map[string]interface{}{"method1": "1", "method2": "2"}))
	log.Error("这个是一起的")
	log.Sync()

	log = u.Core.GetCoreLog().TraceLog("read1")
	log.Error("this is read1", zap.String("read1", "read1"), zap.Any("carry", map[string]interface{}{"method1": "1", "method2": "2"}))
	log.Sync()

	_, _ = u.Core.GetApis().Baidu().GetRoot()

	_, _ = u.Core.GetApis().Taobao().GetRoot()

	u.Core.GetDao().User().CheckPassword("admin", "admin")

	return req.Email
}

func (u *User) FindApi() ([]model.ApiUser, error) {
	return u.Core.GetDao().Native().User.FindApi()
}

func (u *User) Create() {
	panic("implement me")
}

func (u *User) Update() {
	panic("implement me")
}

func (u *User) Delete() {
	panic("implement me")
}
