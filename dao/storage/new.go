package storage

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Storage struct {
	user       types.User
	role       types.Role
	permission types.Permission
	core       types.DaoCore
	db         *gorm.DB
}

func (s *Storage) Permission() types.Permission {
	return s.permission
}

func (s *Storage) Role() types.Role {
	return s.role
}

func (s *Storage) DB() *gorm.DB {
	return s.db.WithContext(s.core.GetTrace().Ctx())
}

func (s *Storage) User() types.User {
	return s.user
}

func NewStorage(db *gorm.DB, enforcer *casbin.Enforcer, core types.DaoCore) types.Storager {
	return &Storage{
		db:         db,
		core:       core,
		user:       NewUser(core, enforcer),
		role:       NewRole(core, enforcer),
		permission: NewPermission(core),
	}
}
