package dbquery

import (
	"database/sql"
	"github.com/fitan/magic/dal/query"
	"github.com/fitan/magic/pkg/types"
)

type WrapQuery struct {
	types.ServiceCore
	*query.Query
}

func (w *WrapQuery) RawQ() *query.Query {
	return w.Query
}

func (w *WrapQuery) WrapQuery() *query.QueryCtx {
	return w.WrapWithContext(w.GetTrace().Ctx())
	//return w.WrapWithContext(w.GetTrace().Ctx())
}

func (w *WrapQuery) Transaction(fc func(tx *query.Query) error, opts ...*sql.TxOptions) error {
	return w.WrapTransaction(w.GetTrace().Ctx())(fc)
}
