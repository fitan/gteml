package types

import "github.com/fitan/magic/model"

type Role interface {
	UnBindPermission(roleID uint, permissionID uint) (err error)
	BindPermission(roleID uint, permissionID uint) (err error)
	Get() (res []model.Role, err error)
	GetById(id uint) (res *model.Role, err error)
	Create(role *model.Role) error
	DeleteById(id uint) (err error)
}
