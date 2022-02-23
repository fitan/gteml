package router

import (
	"github.com/fitan/magic/gen/transfer/k8s"
	"github.com/fitan/magic/gen/transfer/permission"
	"github.com/fitan/magic/gen/transfer/role"
	"github.com/fitan/magic/gen/transfer/user"
	"github.com/fitan/magic/handler/restapi"
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/ginmid"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/prometheus"
	"github.com/fitan/magic/pkg/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	//"go.elastic.co/apm/module/apmgin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Router() *gin.Engine {
	//r := gin.New()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Total-Count"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	r.Use(otelgin.Middleware("ginhttp"))

	//r.Use(apmgin.Middleware(r))
	r.Use(ginmid.SetCore())
	r.Use(ginmid.NewAudit().Audit())

	prometheus.UseGinprom(r)
	pprof.Register(r)
	//r.Use(ginmid.ReUserCore())

	jwtMid, err := ginmid.NewAuthMiddleware()
	if err != nil {
		log.Panicln(err)
	}

	_ = r.Group("/", jwtMid.MiddlewareFunc())
	//registerRouter(authorized)
	registerRouter(r)

	LoginRoute(r, jwtMid)
	//g := r.Group("/api")
	info(r)
	swag(r)
	ping(r)

	db := core.GetCorePool().GetObj().GetDao().Storage().DB()
	userRest := rest.NewBaseRest(db, &restapi.UserObj{})
	roleRest := rest.NewBaseRest(db, &restapi.RolesObj{})
	serviceRest := rest.NewBaseRest(db, &restapi.ServiceObj{})
	rest.RegisterRestApi(r, restapi.NewApiRest(userRest), "/rest/users")
	rest.RegisterRestApi(r, restapi.NewApiRest(roleRest), "/rest/roles")
	rest.RegisterRestApi(r, restapi.NewApiRest(serviceRest), "/rest/services")

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
