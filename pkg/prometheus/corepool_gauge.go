package prometheus

import "github.com/Depado/ginprom"

type CorePoolGauge struct {
	prom *ginprom.Prometheus
	name string
}

func (c *CorePoolGauge) Init(prom *ginprom.Prometheus) {
	c.prom = prom
	c.name = "corepool"
	c.prom.AddCustomGauge(c.name, "corepool status", []string{"method"})
}

func (c *CorePoolGauge) CorePool(method string) {
	c.prom.IncrementGaugeValue(c.name, []string{method})
}
