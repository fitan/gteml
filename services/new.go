package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services/audit"
	"github.com/fitan/magic/services/rbac"
	"github.com/fitan/magic/services/user"
)

type Services struct {
	user  types.Userer
	rabc  types.RBAC
	audit types.Audit
}

func (s *Services) Audit() types.Audit {
	return s.audit
}

func (s *Services) RABC() types.RBAC {
	return s.rabc
}

func (s *Services) User() types.Userer {
	return s.user
}

func NewServices(core types.ServiceCore, enforcer *casbin.Enforcer) types.Serviceser {
	return &Services{
		user.NewUser(core),
		rbac.NewRBAC(enforcer, core),
		audit.NewAudit(core),
	}
}
