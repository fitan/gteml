package dbquery

import (
	"database/sql"
	query2 "github.com/fitan/magic/dao/dal/query"
	"github.com/fitan/magic/pkg/types"
)

type WrapQuery struct {
	types.TracerCore
	*query2.Query
}

func (w *WrapQuery) RawQ() *query2.Query {
	return w.Query
}

func (w *WrapQuery) WrapQuery() *query2.QueryCtx {
	return w.WrapWithContext(w.GetTrace().Ctx())
	//return w.WrapWithContext(w.GetTrace().Ctx())
}

func (w *WrapQuery) Transaction(fc func(tx *query2.Query) error, opts ...*sql.TxOptions) error {
	return w.WrapTransaction(w.GetTrace().Ctx())(fc)
}
