package types

import (
	"github.com/fitan/magic/model"
	"gorm.io/gen/field"
)

type User interface {
	CheckUserPermission(userID uint, serviceID uint, path, method string) (err error)
	CheckPassword(userName string, password string) (*model.User, error)
	ById(id uint, preload ...field.RelationField) (*model.User, error)
	Create(user *model.User) error
	Update(id uint, user *model.User) error
	UnBindPermission(userID, roleID, serviceID uint) (err error)
	BindPermission(userID, roleID, serviceID uint) (err error)
}
