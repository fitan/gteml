package types

import (
	"database/sql"
	query2 "github.com/fitan/magic/dao/dal/query"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Storager interface {
	User() User
	Role() Role
	Permission() Permission
	Query() WrapQuery
	DB() *gorm.DB
}

type WrapQuery interface {
	types.TracerCore
	RawQ() *query2.Query
	WrapQuery() *query2.QueryCtx
	Transaction(fc func(tx *query2.Query) error, opts ...*sql.TxOptions) error
}
