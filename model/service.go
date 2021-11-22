package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string
	Alias       string
	Description string
	ParentId    int64
	Services    []Service `gorm:"foreignkey:ParentId"`
}
