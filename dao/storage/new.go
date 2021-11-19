package storage

import (
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type Storage struct {
	user types.UserModeler
}

func (s *Storage) User() types.UserModeler {
	return s.user
}

func NewStorage(db *gorm.DB) types.Storager {
	return &Storage{
		user: NewUser(db),
	}
}
