package query

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type QueryCtx struct {
	*queryCtx
}

func (q *Query) WrapWithContext(ctx context.Context) *QueryCtx {
	return &QueryCtx{q.WithContext(ctx)}
}

func (q *Query) WrapTransaction(ctx context.Context) func(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return func(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
		return q.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
	}
}
