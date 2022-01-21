package storage

import (
	"github.com/casbin/casbin/v2"
	types2 "github.com/fitan/magic/dao/types"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Storage struct {
	user       types2.User
	role       types2.Role
	permission types2.Permission
	core       types.ServiceCore
	db         *gorm.DB
	query      *dbquery.WrapQuery
}

func (s *Storage) Permission() types2.Permission {
	return s.permission
}

func (s *Storage) Role() types2.Role {
	return s.role
}

func (s *Storage) Query() types.WrapQuery {
	return s.query
}

func (s *Storage) DB() *gorm.DB {
	return s.db.WithContext(s.core.GetTrace().Ctx())
}

func (s *Storage) User() types2.User {
	return s.user
}

func NewStorage(db *gorm.DB, query *dbquery.WrapQuery, enforcer *casbin.Enforcer, core types.ServiceCore) *Storage {
	return &Storage{
		db:         db,
		query:      query,
		core:       core,
		user:       NewUser(query, core, enforcer),
		role:       NewRole(query, core, enforcer),
		permission: NewPermission(query, core),
	}
}
