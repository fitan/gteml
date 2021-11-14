package services

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services/user"
)

type Services struct {
	user types.User
}

func (s *Services) User() types.User {
	return s.user
}

func NewServices(core types.ServiceCore) types.ServicesI {
	return &Services{
		user.NewUser(core),
	}
}
