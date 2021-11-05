package main

import (
	"github.com/fitan/magic/internal/api/router"
	"github.com/fitan/magic/pkg/context"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func main() {
	//tp := trace.GetTp()
	//tr := tp.Tracer("tracer")
	//ctx := context.Background()
	//spanctx, _ := tr.Start(ctx, "log1")
	//span := trace2.SpanFromContext(spanctx)
	//span.AddEvent(semconv.ExceptionEventName, trace2.WithAttributes(semconv.ExceptionTypeKey.String("log"), semconv.ExceptionMessageKey.String(string("this is log 1"))))
	//span.RecordError(fmt.Errorf("this is error %s", "log1"))
	//span.SetStatus(1, "statuso")
	//span.Sync()
	if context.Conf.MyConf.Pyroscope.Open {
		profiler.Start(
			profiler.Config{
				ApplicationName: context.GetCtxPool().GetObj().Config.App.Name,

				// replace this with the address of pyroscope server
				ServerAddress: context.Conf.MyConf.Pyroscope.Url,

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

	router.Router().Run()
}
