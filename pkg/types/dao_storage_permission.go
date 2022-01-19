package types

import "github.com/fitan/magic/dao/dal/model"

type Permission interface {
	Create(permission *model.Permission) (err error)
	Get() (res []*model.Permission, err error)
	GetByID(id uint) (res *model.Permission, err error)
	DeleteByID(id uint) (err error)
	UpdateById(permission *model.Permission) (err error)
}
