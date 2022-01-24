package main

import (
	"fmt"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/fitan/magic/pkg/core"
	micro2 "github.com/fitan/magic/pkg/micro"
	"github.com/fitan/magic/router"
	"github.com/gin-gonic/gin"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"log"
	"os"
	"time"
)

var (
	gitHash   string
	gitTag    string
	buildTime string
	goVersion string
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {

	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Tag: %s \n", gitTag)
		fmt.Printf("Git Commit hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
		return
	}
	core.NewCore()

	if core.ConfReg.Confer.GetMyConf().Pyroscope.Open {
		profiler.Start(
			profiler.Config{
				ApplicationName: core.ConfReg.Confer.GetMyConf().App.Name,

				// replace this with the address of pyroscope server
				ServerAddress: core.ConfReg.Confer.GetMyConf().Pyroscope.Url,

				// by default all profilers are enabled,
				// but you can select the ones you want to use:
				ProfileTypes: []profiler.ProfileType{
					profiler.ProfileCPU,
					profiler.ProfileAllocObjects,
					profiler.ProfileAllocSpace,
					profiler.ProfileInuseObjects,
					profiler.ProfileInuseSpace,
				},
			},
		)
	}

	r := router.Router()

	srv := httpServer.NewServer(
		server.Name("gteml"),
		server.Address(":8080"),
	)
	gin.SetMode(gin.ReleaseMode)

	hd := srv.NewHandler(r)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(micro2.ConsulRegistry(core.ConfReg.Confer.GetMyConf().Consul.Addr)),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()
	err := service.Run()
	if err != nil {
		log.Printf("mircro run error %v\n", err)
	}

}
