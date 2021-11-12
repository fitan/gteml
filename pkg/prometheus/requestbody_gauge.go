package prometheus

import (
	"fmt"
	"github.com/Depado/ginprom"
)

type RequestBodyGauge struct {
	prom *ginprom.Prometheus
	name string
}

func (r *RequestBodyGauge) Init(prom *ginprom.Prometheus) {
	r.prom = prom
	r.name = "request_body"
	//r.name = strings.Join([]string{r.prom.Namespace, r.prom.Subsystem, "request_body"}, "_")
	r.prom.AddCustomGauge(r.name, "body internal state code", []string{"code", "method", "path"})
}

func (r *RequestBodyGauge) RequestBody(code string, method string, path string) {
	err := r.prom.IncrementGaugeValue(r.name, []string{code, method, path})
	if err != nil {
		fmt.Println(err.Error())
	}
}
