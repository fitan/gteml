package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services/rbac"
	"github.com/fitan/magic/services/user"
)

type Services struct {
	user types.Userer
	rabc casbin.IEnforcer
}

func (s *Services) RABC() casbin.IEnforcer {
	return s.rabc
}

func (s *Services) User() types.Userer {
	return s.user
}

func NewServices(core types.ServiceCore, enforcer *casbin.Enforcer) types.Serviceser {
	return &Services{
		user.NewUser(core),
		rbac.NewRBAC(enforcer),
	}
}
