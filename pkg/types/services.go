package types

import (
	"github.com/fitan/magic/services/types"
)

type Serviceser interface {
	User() types.Userer
	RABC() types.RBAC
	Audit() types.Audit
	K8s() types.K8s
}
