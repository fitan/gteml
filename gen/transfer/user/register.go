package user

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *ginx.GinXHandlerRegister) {

	reg.Register(r, &BindUserPermissionTransfer{}, ginx.WithHandlerName("BindUserPermission"))

	reg.Register(r, &UnBindUserPermissionTransfer{}, ginx.WithHandlerName("UnBindUserPermission"))

	reg.Register(r, &GetUserByIDTransfer{}, ginx.WithHandlerName("GetUserByID"))

	reg.Register(r, &CreateTransfer{}, ginx.WithHandlerName("Create"))

	reg.Register(r, &SayHelloTransfer{}, ginx.WithHandlerName("SayHello"))

}
