package types

import (
	"github.com/fitan/magic/model"
)

type Serviceser interface {
	User() Userer
	RABC() RBAC
}

type Userer interface {
	Create()
	Update()
	Delete()
	Read() string
	Login(username string, password string) (*model.User, error)
}
