package role

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *ginx.GinXHandlerRegister) {

	reg.Register(r, &BindRolePermissionTransfer{}, ginx.WithHandlerName("BindRolePermission"))

	reg.Register(r, &UnBindRolePermissionTransfer{}, ginx.WithHandlerName("UnBindRolePermission"))

}
