package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/broker/redis/v4"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/fitan/magic/pkg/core"
	micro2 "github.com/fitan/magic/pkg/micro"
	"github.com/fitan/magic/router"
	"github.com/gin-gonic/gin"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"log"
	"os"
	"os/signal"
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
		server.Address(":8081"),
	)
	gin.SetMode(gin.ReleaseMode)

	hd := srv.NewHandler(r)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt, os.Kill)

	reg := micro2.EtcdRegistry(core.ConfReg.Confer.GetMyConf().Consul.Addr)
	service := micro.NewService(
		micro.Context(ctx),
		micro.Server(srv),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.HandleSignal(false),
		micro.Broker(redis.NewBroker(broker.Addrs("localhost:6379"))),
	)
	//defaultLock := "micro-sync-lock"
	//consulSync := consul.NewSync(sync.Nodes(core.ConfReg.Confer.GetMyConf().Consul.Addr))

	go func() {
		//err := consulSync.Lock(defaultLock)
		//if err != nil {
		//	fmt.Printf("lock err: %v", err)
		//}
		//fmt.Println("lock ok")
		//leader, err := consulSync.Leader(defaultLock)
		//if err != nil {
		//	fmt.Printf("leaderC")
		//}
		//go func() {
		//	status := leader.Status()
		//	for {
		//		s := <-status
		//		fmt.Printf("status: %v", s)
		//	}
		//}()

		<-sc
		//err = consulSync.Unlock(defaultLock)
		//if err != nil {
		//	fmt.Printf("unlock err %v", err)
		//}
		//fmt.Println("unlock ok")
		id := service.Options().Server.Options().Name + "-" + service.Options().Server.Options().Id
		err := reg.Deregister(
			&registry.Service{
				Name:    service.Options().Server.Options().Name,
				Version: service.Options().Server.Options().Version,
				Nodes:   []*registry.Node{&registry.Node{Id: id}},
			},
		)
		if err != nil {
			log.Printf("Deregisterr %v. err: %v\n", id, err)
		} else {
			log.Printf("Deregistering node: %v\n", id)
		}
		<-time.After(time.Second * 10)
		cancel()
	}()

	service.Init()
	err := service.Run()
	if err != nil {
		log.Printf("mircro run error %v\n", err)
	}

}
