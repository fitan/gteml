package types

import "github.com/fitan/magic/model"

type User interface {
	CheckUserPermission(userID uint, serviceID uint, path, method string) error
	CheckPassword(userName string, password string) (*model.User, error)
	ById(id uint, preload ...string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	UnBindPermission(userID, roleID, serviceID uint) (err error)
	BindPermission(userID, roleID, serviceID uint) (err error)
}
