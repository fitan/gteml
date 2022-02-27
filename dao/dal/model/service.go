package model

type Service struct {
	Model
	Name        string
	Alias       string
	Description string
	ParentId    uint
	Services    []Service `gorm:"foreignKey:ParentId"`
}
