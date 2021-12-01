package storage

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Storage struct {
	user       types.User
	role       types.Role
	permission types.Permission
	core       types.ServiceCore
	db         *gorm.DB
	query      *dbquery.WrapQuery
}

func (s *Storage) Permission() types.Permission {
	return s.permission
}

func (s *Storage) Role() types.Role {
	return s.role
}

func (s *Storage) Query() types.WrapQuery {
	return s.query
}

func (s *Storage) DB() *gorm.DB {
	return s.db.WithContext(s.core.GetTrace().Ctx())
}

func (s *Storage) User() types.User {
	return s.user
}

func NewStorage(db *gorm.DB, query *dbquery.WrapQuery, enforcer *casbin.Enforcer, core types.ServiceCore) types.Storager {
	return &Storage{
		db:         db,
		query:      query,
		core:       core,
		user:       NewUser(query, core, enforcer),
		role:       NewRole(query, core, enforcer),
		permission: NewPermission(query, core),
	}
}
