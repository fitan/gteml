package user

import "github.com/fitan/magic/pkg/types"

func NewUser(core types.ServiceCore) types.User {
	return &User{core}
}

type User struct {
	Core types.ServiceCore
}

func (u *User) Read() string {
	log := u.Core.GetCoreLog().TraceLog("user.read")
	defer log.Sync()

	return u.Core.GetConfig().GetMyConf().App.Name
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
