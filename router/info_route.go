package router

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func info(i gin.IRouter) {
	i.GET("/routers", func(c *gin.Context) {
		ip, _ := c.RemoteIP()
		if ip.String() == "::1" {
			c.JSON(http.StatusOK, ginx.CollectRouterSlice)
			return
		}
		c.JSON(http.StatusOK, ip.String())
	})
}
