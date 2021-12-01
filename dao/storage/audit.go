package storage

import (
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/dbquery"
)

type Audit struct {
	query *dbquery.WrapQuery
}

func (a *Audit) Create(audit *model.Audit) error {
	return a.query.WrapQuery().Audit.Create(audit)
}
