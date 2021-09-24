package httpclient

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func ErrorTrace(tr trace.Tracer) resty.ErrorHook {
	return func(request *resty.Request, err error) {
		traceInfo := new(TraceInfo)
		if v, ok := err.(*resty.ResponseError); ok {
			traceInfo.Response = SetResponse(v.Response)
		}
		traceInfo.Request = SetRequest(request)
		traceInfo.Info = SetInfo(request.TraceInfo())
		traceRaw, _ := json.Marshal(traceInfo)
		subContext, span := tr.Start(request.Context(), "error_hook")
		span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("info"), semconv.ExceptionMessageKey.String(string(traceRaw))))
		span.SetStatus(codes.Error, "error_hook")
		span.End()
		context.WithValue(request.Context(), "sub_ctx", subContext)
	}
}

func BeforeTrace(tp trace.TracerProvider) resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		client.SetTransport(otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp)))
		return nil
	}
}

func AfterTraceDebug(tr trace.Tracer) resty.ResponseMiddleware {
	return func(client *resty.Client, response *resty.Response) error {
		traceInfo := new(TraceInfo)
		traceInfo.Request = SetRequest(response.Request)
		traceInfo.Response = SetResponse(response)
		traceInfo.Info = SetInfo(response.Request.TraceInfo())
		traceRaw, _ := json.Marshal(traceInfo)
		subContext, span := tr.Start(response.Request.Context(), "trace_debug")
		span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("info"), semconv.ExceptionMessageKey.String(string(traceRaw))))
		span.End()
		context.WithValue(response.Request.Context(), "sub_ctx", subContext)
		return nil
	}
}

type DynamicHostHooker interface {
	GetHost() string
}

func DynamicHostHook(hooker DynamicHostHooker) resty.PreRequestHook {
	return func(client *resty.Client, request *http.Request) error {
		request.Host = hooker.GetHost()
		return nil
	}
}

//func BeforSelectNode() resty.RequestMiddleware {
//	return func(client *resty.Client, request *resty.Request) error {
//	}
//}
