package rbac

import (
	"github.com/casbin/casbin/v2"
)

func NewRBAC(enforcer *casbin.Enforcer) *RBAC {
	return &RBAC{
		enforcer,
	}

}

type RBAC struct {
	*casbin.Enforcer
}
