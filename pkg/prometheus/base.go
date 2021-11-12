package prometheus

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

var g *ginprom.Prometheus

func GetGinprom() *ginprom.Prometheus {
	if g == nil {
		g = ginprom.New(
			ginprom.Path("/metrics"),
		)
	}

	return g
}

func UseGinprom(r *gin.Engine) *ginprom.Prometheus {
	g := GetGinprom()
	g.Use(r)
	r.Use(g.Instrument())
	return g
}
