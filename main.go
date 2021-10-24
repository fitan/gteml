package main

import (
	"github.com/fitan/gteml/internal/api/router"
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

	router.Router().Run()
}
