package types

import (
	"database/sql"
	query2 "github.com/fitan/magic/dao/dal/query"
	"github.com/fitan/magic/dao/types"
	"gorm.io/gorm"
)

type DAOer interface {
	Storage() Storager
}

type Storager interface {
	User() types.User
	Role() types.Role
	Permission() types.Permission
	Query() WrapQuery
	DB() *gorm.DB
}

type WrapQuery interface {
	TracerCore
	RawQ() *query2.Query
	WrapQuery() *query2.QueryCtx
	Transaction(fc func(tx *query2.Query) error, opts ...*sql.TxOptions) error
}
