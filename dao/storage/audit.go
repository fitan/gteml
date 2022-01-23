package storage

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/dbquery"
)

type Audit struct {
	query *dbquery.WrapQuery
}

func NewAudit(query *dbquery.WrapQuery) *Audit {
	return &Audit{query: query}
}

func (a *Audit) Find(key string, page, pageSize int) ([]*model.Audit, int64, error) {
	q := a.query
	do := a.query.WrapQuery().Audit.Where()

	if key != "" {
		key = "%" + key + "%"
		do = do.Where(
			q.RawQ().Audit.Method.Like(key),
			q.RawQ().Audit.Request.Like(key),
			q.RawQ().Audit.Response.Like(key),
			q.RawQ().Audit.Query.Like(key),
			q.RawQ().Audit.Header.Like(key),
			q.RawQ().Audit.Url.Like(key),
		)
	}

	var offset int
	if !(page > 0 && pageSize > 0) {
		offset = -1
	} else {
		offset = (page - 1) * pageSize
	}

	return do.FindByPage(offset, pageSize)
}
