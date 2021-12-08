package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services/audit"
	"github.com/fitan/magic/services/k8s"
	"github.com/fitan/magic/services/rbac"
	"github.com/fitan/magic/services/user"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Services struct {
	user  types.Userer
	rabc  types.RBAC
	audit types.Audit
	k8s   types.K8s
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

func (s *Services) K8s() types.K8s {
	return s.k8s
}

func NewServices(core types.ServiceCore, enforcer *casbin.Enforcer, k8sClient *kubernetes.Clientset, runtimeClient client.Client) types.Serviceser {
	return &Services{
		user.NewUser(core),
		rbac.NewRBAC(enforcer, core),
		audit.NewAudit(core),
		k8s.NewK8s(k8sClient, runtimeClient, core),
	}
}
