package main

import (
	"fmt"
	"github.com/fitan/magic/pkg/core"
	router2 "github.com/fitan/magic/router"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"os"
)

var (
	gitHash   string
	gitTag    string
	buildTime string
	goVersion string
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Tag: %s \n", gitTag)
		fmt.Printf("Git Commit hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
		return
	}
	//tp := trace.GetTp()
	//tr := tp.Tracer("tracer")
	//ctx := context.Background()
	//spanctx, _ := tr.Start(ctx, "log1")
	//span := trace2.SpanFromContext(spanctx)
	//span.AddEvent(semconv.ExceptionEventName, trace2.WithAttributes(semconv.ExceptionTypeKey.String("log"), semconv.ExceptionMessageKey.String(string("this is log 1"))))
	//span.RecordError(fmt.Errorf("this is error %s", "log1"))
	//span.SetStatus(1, "statuso")
	//span.Sync()
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

	router2.Router().Run()
}
