package storage

import (
	"github.com/fitan/magic/pkg/ent"
	"github.com/fitan/magic/pkg/types"
)

func NewStorage(ctx *types.Context, client *ent.Client) types.Storage {
	return &Storage{
		NewUser(ctx, client),
	}
}

type Storage struct {
	user types.UserInterface
}

func (s *Storage) User() types.UserInterface {
	return s.user
}
