package model

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	Url      string
	Method   string
	UserName string
	UserID   int64
	Body     string
}
