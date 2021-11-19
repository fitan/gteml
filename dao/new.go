package dao

import (
	"github.com/fitan/magic/dao/storage"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

type DAO struct {
	storage types.Storager
}

func (d *DAO) Storage() types.Storager {
	return d.storage
}

func NewDAO(db *gorm.DB) *DAO {
	return &DAO{
		storage: storage.NewStorage(db),
	}
}
