package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"log"
)

var (
	tp    *trace.TracerProvider
	stdTp *trace.TracerProvider
)

func StdTracerProvider() {
	var err error
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
		return
	}
	bsp := trace.NewBatchSpanProcessor(exp)
	stdTp = trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(stdTp)
}

func TracerProvider(serviceName string, url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func Init() {
	var err error
	tp, err = TracerProvider("demo", "http://10.170.34.122:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	//tr := tp.Tracer("cmdb")
	//ctx, _ := context.WithCancel(context.Background())
	//ctx, span := tr.Start(ctx, "foo")
	//defer span.Sync()
	//tr = otel.Tracer("new-")
	//ctx, span = tr.Start(ctx, "bar")
	//defer span.Sync()
}

func GetTp() *trace.TracerProvider {
	return tp
}

func GetTrCxt() context.Context {
	StdTracerProvider()
	tr := tp.Tracer("cmdb")
	ctx := context.Background()
	ctx, span := tr.Start(ctx, "ent")
	defer span.End()
	tr = tp.Tracer("new_cmdb")
	ctx, span = tr.Start(ctx, "fsdf")
	span.End()
	return ctx
}

func GetTr() oteltrace.Tracer {
	tr := tp.Tracer("cmdb")
	ctx := context.Background()
	ctx, span := tr.Start(ctx, "ent")
	defer span.End()
	return tp.Tracer("new_cmdb")
}
