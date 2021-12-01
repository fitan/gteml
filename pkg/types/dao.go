package types

import (
	"database/sql"
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
	Query() WrapQuery
	DB() *gorm.DB
}

type WrapQuery interface {
	ServiceCore
	RawQ() *query.Query
	WrapQuery() *query.QueryCtx
	Transaction(fc func(tx *query.Query) error, opts ...*sql.TxOptions) error
}
