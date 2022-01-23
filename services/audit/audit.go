package audit

import (
	"github.com/fitan/magic/dao/dal/model"
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
	return a.core.GetDao().Storage().Native().Audit.Create(audit)
}

func (a *Audit) Find(key string, page, pageSize int) ([]*model.Audit, int64, error) {
	return a.core.GetDao().Storage().Audit().Find(key, page, pageSize)
}
