package audit

import (
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
)

type Audit struct {
	core types.ServiceCore
}

func NewAudit(core types.ServiceCore) *Audit {
	return &Audit{core: core}
}

func (a *Audit) InsetAudit(audit *model.Audit) error {
	log := a.core.GetCoreLog().TraceLog("services.audit.insetAudit")
	defer func() {
		log.Sync()
	}()
	return a.core.GetDao().Storage().Query().WrapQuery().Audit.Create(audit)
}

func (a *Audit) Find(key string, page, pageSize int) ([]*model.Audit, int64, error) {
	q := a.core.GetDao().Storage().Query()
	do := q.WrapQuery().Audit.Where()

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
