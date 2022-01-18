package main

import (
	"fmt"
	"github.com/fitan/magic/pkg/core"
	router2 "github.com/fitan/magic/router"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"log"
	"os"
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

	err := router2.Router().Run()
	log.Printf("gin run error %v\n", err)
}
