package types

import (
	"github.com/fitan/magic/dal/query"
	"gorm.io/gorm"
)

type DAOer interface {
	Storage() Storager
}

type Storager interface {
	User() User
	Role() Role
	Permission() Permission
	Query() *query.WrapQuery
	DB() *gorm.DB
}
