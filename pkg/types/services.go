package types

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/model"
)

type Serviceser interface {
	User() Userer
	RABC() casbin.IEnforcer
}

type Userer interface {
	Create()
	Update()
	Delete()
	Read() string
	Login(username string, password string) (*model.User, error)
}
