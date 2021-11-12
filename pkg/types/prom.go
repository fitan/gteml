package types

import "github.com/Depado/ginprom"

type Promer interface {
	CorePool(method string)
	RequestBody(code string, method string, path string)
}

type InitGauger interface {
	Init(prom *ginprom.Prometheus)
}
