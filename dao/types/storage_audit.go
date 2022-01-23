package types

import "github.com/fitan/magic/dao/dal/model"

type Audit interface {
	Find(key string, page, pageSize int) ([]*model.Audit, int64, error)
}
