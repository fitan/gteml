package types

import "github.com/fitan/magic/dao/dal/model"

type Audit interface {
	InsetAudit(audit *model.Audit) error
	Find(key string, page, pageSize int) ([]*model.Audit, int64, error)
}
