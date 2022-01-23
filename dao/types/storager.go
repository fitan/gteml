package types

import (
	"github.com/fitan/magic/dao/dal/query"
)

type Storager interface {
	User() User
	Role() Role
	Permission() Permission
	Audit() Audit
	Native() *query.QueryCtx
}
