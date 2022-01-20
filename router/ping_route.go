package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(i gin.IRouter) {
	i.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
}
