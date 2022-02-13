package types

import (
	"github.com/fitan/magic/dao/dal/query"
	"gorm.io/gorm"
)

type Storager interface {
	DB() *gorm.DB
	User() User
	Role() Role
	Permission() Permission
	Audit() Audit
	Native() *query.QueryCtx
}
