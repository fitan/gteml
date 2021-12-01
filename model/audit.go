package model

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	Url        string
	Query      string
	Method     string
	Request    string
	Response   string
	Header     string
	StatusCode int
	RemoteIP   string
	ClientIP   string
	CostTime   string
}
