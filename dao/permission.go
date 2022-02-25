package dao

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
)

type Permission struct {
	query *dbquery.WrapQuery
	//core types.ServiceCore
}

func NewPermission(query *dbquery.WrapQuery, core types.ServiceCore) *Permission {
	return &Permission{query: query}
}

func (p *Permission) Create(permission *model.Permission) (err error) {
	return p.query.WrapQuery().Permission.Create(permission)
}

func (p *Permission) UpdateById(permission *model.Permission) (err error) {
	_, err = p.query.WrapQuery().Permission.Where().Updates(permission)
	return
}

func (p *Permission) Get() (res []*model.Permission, err error) {
	return p.query.WrapQuery().Permission.Find()
}

func (p *Permission) GetByID(id uint) (res *model.Permission, err error) {
	return p.query.WrapQuery().Permission.Where(p.query.Permission.ID.Eq(id)).First()
}

func (p *Permission) DeleteByID(id uint) (err error) {
	_, err = p.query.WrapQuery().Permission.Where(p.query.Permission.ID.Eq(id)).Delete()
	return
}
