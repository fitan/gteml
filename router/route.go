package router

import (
	"github.com/fitan/magic/gen/transfer/k8s"
	"github.com/fitan/magic/gen/transfer/permission"
	"github.com/fitan/magic/gen/transfer/role"
	"github.com/fitan/magic/gen/transfer/user"
	"github.com/fitan/magic/pkg/ginmid"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/prometheus"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
	"log"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(apmgin.Middleware(r))
	r.Use(ginmid.SetCore())
	r.Use(ginmid.NewAudit().Audit())

	prometheus.UseGinprom(r)
	pprof.Register(r)
	//r.Use(ginmid.ReUserCore())

	jwtMid, err := ginmid.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	authorized := r.Group("/", jwtMid.MiddlewareFunc())
	registerRouter(authorized)
	LoginRoute(r, jwtMid)
	SwagRoute(r)

	return r
}

// 注册路由
func registerRouter(r gin.IRouter) {
	gReg := ginx.NewGinXHandlerRegister()
	role.Register(r, gReg)
	user.Register(r, gReg)
	k8s.Register(r, gReg)
	permission.Register(r, gReg)
}
