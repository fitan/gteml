package storage

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/dal/query"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
)

type Role struct {
	core     types.ServiceCore
	query    *dbquery.WrapQuery
	enforcer *casbin.Enforcer
}

func NewRole(query *dbquery.WrapQuery, core types.ServiceCore, enforcer *casbin.Enforcer) *Role {
	return &Role{query: query, core: core, enforcer: enforcer}
}

func (r *Role) UnBindPermission(roleID uint, permissionID uint) (err error) {

	err = r.query.Transaction(func(tx *query.Query) error {
		role, err := tx.Role.Where(tx.Role.ID.Eq(roleID)).First()
		if err != nil {
			return err
		}

		permission, err := tx.Permission.Where(tx.Permission.ID.Eq(permissionID)).First()
		if err != nil {
			return err
		}
		_, err = r.enforcer.RemovePolicy(roleID2CasbinKey(roleID), permission.Path, permission.Method)
		if err != nil {
			return err
		}

		err = tx.Role.Permissions.Model(role).Delete(permission)
		if err != nil {
			return err
		}

		return nil

	})

	return
}

func (r *Role) BindPermission(roleID uint, permissionID uint) (err error) {

	err = r.query.Transaction(func(tx *query.Query) error {
		role, err := tx.Role.Where(tx.Role.ID.Eq(roleID)).First()
		if err != nil {
			return err
		}

		permission, err := tx.Permission.Where(tx.Permission.ID.Eq(permissionID)).First()
		if err != nil {
			return err
		}

		_, err = r.enforcer.AddPolicy(roleID2CasbinKey(roleID), permission.Path, permission.Method)
		if err != nil {
			return err
		}

		err = tx.Role.Permissions.Model(role).Append(permission)
		if err != nil {
			return err
		}
		return nil
	})

	return
}

func (r *Role) Get() (res []*model.Role, err error) {
	return r.query.Role.Find()
}

func (r *Role) GetById(id uint) (res *model.Role, err error) {
	return r.query.WrapQuery().Role.Where(r.query.Role.ID.Eq(id)).First()
}

func (r *Role) Create(role *model.Role) error {
	return r.query.WrapQuery().Role.Create(role)
}

func (r *Role) DeleteById(id uint) (err error) {

	err = r.query.Transaction(func(tx *query.Query) error {
		role, err := tx.Role.Where(tx.Role.ID.Eq(id)).First()
		if err != nil {
			return err
		}

		err = tx.Role.Permissions.Model(role).Clear()
		if err != nil {
			return err
		}

		_, err = r.enforcer.RemoveFilteredPolicy(0, role.OnlyKey)
		if err != nil {
			return err
		}

		_, err = tx.Role.Where(tx.Role.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}

		return nil
	})

	return
}
