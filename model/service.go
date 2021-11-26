package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string
	Alias       string
	Description string
	ParentId    uint
	Services    []Service `gorm:"foreignKey:ParentId"`
}
