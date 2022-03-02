package prometheus

import (
	"github.com/Depado/ginprom"
	"github.com/fitan/magic/pkg/types"
	"log"
	"reflect"
)

type Gauge struct {
	*RequestBodyGauge
	*CorePoolGauge
	*HttpclientGauge
}

func NewGauge(prom *ginprom.Prometheus) types.Promer {
	g := &Gauge{
		RequestBodyGauge: &RequestBodyGauge{},
		CorePoolGauge:    &CorePoolGauge{},
		HttpclientGauge:  &HttpclientGauge{},
	}

	rv := reflect.ValueOf(g)
	for i := 0; i < rv.Elem().NumField(); i++ {
		f := rv.Elem().Field(i)

		v, ok := f.Interface().(types.InitGauger)
		if !ok {
			log.Panicf("field %s not implemented InitGauger")
		}

		v.Init(prom)
	}
	return g
}
