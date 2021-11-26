package storage

import (
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
)

type Permission struct {
	core types.DaoCore
}

func NewPermission(core types.DaoCore) *Permission {
	return &Permission{core: core}
}

func (p *Permission) Create(permission *model.Permission) (err error) {
	db := p.core.GetDao().Storage().DB()

	err = db.Create(permission).Error
	return
}

func (p *Permission) UpdateById(permission *model.Permission) (err error) {
	db := p.core.GetDao().Storage().DB()

	err = db.Save(permission).Error
	return
}

func (p *Permission) Get() (res []model.Permission, err error) {
	db := p.core.GetDao().Storage().DB()
	err = db.Find(&res).Error
	return
}

func (p *Permission) GetByID(id uint) (res model.Permission, err error) {
	db := p.core.GetDao().Storage().DB()
	err = db.First(&res, id).Error
	return
}

func (p *Permission) DeleteByID(id uint) (err error) {
	db := p.core.GetDao().Storage().DB()
	err = db.Delete(&model.Permission{}, id).Error
	return
}
