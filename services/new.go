package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services/audit"
	"github.com/fitan/magic/services/k8s"
	"github.com/fitan/magic/services/rbac"
	types2 "github.com/fitan/magic/services/types"
	"github.com/fitan/magic/services/user"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Services struct {
	user  types2.Userer
	rabc  types2.RBAC
	audit types2.Audit
	k8s   types2.K8s
}

func (s *Services) Audit() types2.Audit {
	return s.audit
}

func (s *Services) RABC() types2.RBAC {
	return s.rabc
}

func (s *Services) User() types2.Userer {
	return s.user
}

func (s *Services) K8s() types2.K8s {
	return s.k8s
}

func NewServices(core types.ServiceCore, enforcer *casbin.Enforcer, k8sClient *kubernetes.Clientset, runtimeClient client.Client, cfg *rest.Config) types.Serviceser {
	return &Services{
		user.NewUser(core),
		rbac.NewRBAC(enforcer, core),
		audit.NewAudit(core),
		k8s.NewK8s(k8sClient, runtimeClient, cfg, core),
	}
}
