package permission

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *core.GinXHandlerRegister) {

	reg.Register(r, &CreatePermissionTransfer{}, ginx.WithHandlerName("CreatePermission"))

	reg.Register(r, &GetPermissionByIdTransfer{}, ginx.WithHandlerName("GetPermissionById"))

	reg.Register(r, &DeletePermissionByIdTransfer{}, ginx.WithHandlerName("DeletePermissionById"))

	reg.Register(r, &UpdatePermissionTransfer{}, ginx.WithHandlerName("UpdatePermission"))

}
