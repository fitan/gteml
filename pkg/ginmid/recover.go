package ginmid

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Recover() gin.HandlerFunc {
	return gin.CustomRecovery(
		func(c *gin.Context, err interface{}) {

			key, _ := c.Get(types.CoreKey)
			core := key.(*types.Core)

			core.GinX.SetError(errors.New("系统错误,请联系管理员"))
			log := core.GetCoreLog().ApmLog("ginmid.recover")
			log.Error("recover", zap.Any("err", err))
			log.Sync()

		})
}
