package model

type Audit struct {
	Model
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
