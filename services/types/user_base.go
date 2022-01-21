package types

import "github.com/fitan/magic/dao/dal/model"

type Userer interface {
	Create()
	Update()
	Delete()
	Read() string
	Login(username string, password string) (*model.User, error)
}
