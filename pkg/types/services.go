package types

import (
	"github.com/fitan/magic/dao/dal/model"
)

type Serviceser interface {
	User() Userer
	RABC() RBAC
	Audit() Audit
	K8s() K8s
}

type Userer interface {
	Create()
	Update()
	Delete()
	Read() string
	Login(username string, password string) (*model.User, error)
}
