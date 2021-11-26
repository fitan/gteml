package storage

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Role struct {
	core     types.DaoCore
	enforcer *casbin.Enforcer
}

func NewRole(core types.DaoCore, enforcer *casbin.Enforcer) *Role {
	return &Role{core: core, enforcer: enforcer}
}

func (r *Role) UnBindPermission(roleID uint, permissionID uint) (err error) {
	db := r.core.GetDao().Storage().DB()

	err = db.Transaction(
		func(tx *gorm.DB) error {
			role := &model.Role{}
			err = tx.First(role, roleID).Error
			if err != nil {
				return err
			}

			permission := &model.Permission{}
			err = tx.First(permission, permissionID).Error
			if err != nil {
				return err
			}

			_, err = r.enforcer.RemovePolicy(roleID2CasbinKey(roleID), permission.Path, permission.Method)
			if err != nil {
				return err
			}

			err = tx.Model(role).Association("Permissions").Delete(permission)
			if err != nil {
				return err
			}

			return nil
		})
	return
}

func (r *Role) BindPermission(roleID uint, permissionID uint) (err error) {
	db := r.core.GetDao().Storage().DB()

	err = db.Transaction(
		func(tx *gorm.DB) error {
			role := &model.Role{}
			err = tx.First(role, roleID).Error
			if err != nil {
				return err
			}

			permission := &model.Permission{}
			err = tx.First(permission, permissionID).Error
			if err != nil {
				return err
			}

			_, err = r.enforcer.AddPolicy(roleID2CasbinKey(roleID), permission.Path, permission.Method)
			if err != nil {
				return err
			}

			err = tx.Model(role).Association("Permissions").Append(permission)
			if err != nil {
				return err
			}

			return nil
		})
	return
}

func (r *Role) Get() (res []model.Role, err error) {
	db := r.core.GetDao().Storage().DB()
	err = db.Find(&res).Error
	return
}

func (r *Role) GetById(id uint) (res *model.Role, err error) {
	db := r.core.GetDao().Storage().DB()
	err = db.First(res, id).Error
	return
}

func (r *Role) Create(role *model.Role) error {
	db := r.core.GetDao().Storage().DB()
	return db.Create(role).Error
}

func (r *Role) DeleteById(id uint) (err error) {
	db := r.core.GetDao().Storage().DB()
	db.Transaction(
		func(tx *gorm.DB) error {
			role := &model.Role{}
			role.ID = id
			err = tx.Model(role).Association("Permissions").Clear()
			if err != nil {
				return err
			}

			_, err = r.enforcer.RemoveFilteredPolicy(0, role.OnlyKey)
			if err != nil {
				return err
			}

			err = tx.Delete(role).Error
			if err != nil {
				return err
			}

			return nil

		})
	return
}
