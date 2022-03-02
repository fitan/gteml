package prometheus

import (
	"fmt"
	"github.com/Depado/ginprom"
)

type HttpclientGauge struct {
	prom *ginprom.Prometheus
	name string
}

func (h *HttpclientGauge) Init(prom *ginprom.Prometheus) {
	h.prom = prom
	h.name = "httpclient"
	h.prom.AddCustomGauge(h.name+"_status", "http client status", []string{"domain", "code"})
	h.prom.AddCustomGauge(h.name+"_err", "http client error", []string{"domain"})
}

func (h *HttpclientGauge) ClientStatus(domain string, code string) {
	err := h.prom.IncrementGaugeValue(h.name+"_status", []string{domain, code})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (h *HttpclientGauge) ClientErr(domain string) {
	err := h.prom.IncrementGaugeValue(h.name+"_err", []string{domain})
	if err != nil {
		fmt.Println(err.Error())
	}
}
