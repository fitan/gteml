package query

import (
	"database/sql"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type WrapQuery struct {
	*types.Core
	*Query
}

func (w *WrapQuery) WrapQuery() *queryCtx {
	return w.WithContext(w.GetTrace().Ctx())
}

func (w *WrapQuery) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return w.db.WithContext(w.GetTrace().Ctx()).Transaction(func(tx *gorm.DB) error { return fc(w.clone(tx)) }, opts...)
}
