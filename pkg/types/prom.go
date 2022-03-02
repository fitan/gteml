package types

import "github.com/Depado/ginprom"

type Promer interface {
	CorePool(method string)
	RequestBody(code string, method string, path string)
	ClientStatus(domain string, code string)
	ClientErr(domain string)
}

type InitGauger interface {
	Init(prom *ginprom.Prometheus)
}
