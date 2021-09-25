package router

import (
	"github.com/fitan/gteml/internal/api/gen/transfer/user"
	"github.com/fitan/gteml/pkg/core"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	core.GinXHandlerRegister(r, &user.CreateTransfer{})

	return r
}
