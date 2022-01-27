package httpclient

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/registry/cache"
	"go-micro.dev/v4/selector"
	"go.elastic.co/apm/module/apmhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func parseCtxValOpenTrace(ctx context.Context) bool {
	return ctx.Value(_OpenTrace).(bool)
}

func AfterErrorTrace() resty.ErrorHook {
	return func(request *resty.Request, err error) {
		if !parseCtxValOpenTrace(request.Context()) {
			return
		}
		traceInfo := new(TraceInfo)
		if v, ok := err.(*resty.ResponseError); ok {
			traceInfo.Response = SetResponse(v.Response)
		}
		traceInfo.Request = SetRequest(request)
		traceInfo.Info = SetInfo(request.TraceInfo())
		traceRaw, _ := json.Marshal(traceInfo)
		span := trace.SpanFromContext(request.Context())
		span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("info"), semconv.ExceptionMessageKey.String(string(traceRaw))))
		span.SetStatus(codes.Error, "error_hook")
		span.End()
		span = trace.SpanFromContext(request.Context())
	}
}

// 当没有触发error时不会触发span end。在这里处理
func AfterErrorSpanEnd() resty.ResponseMiddleware {
	return func(client *resty.Client, response *resty.Response) error {
		if !parseCtxValOpenTrace(response.Request.Context()) {
			return nil
		}
		span := trace.SpanFromContext(response.Request.Context())
		if span.IsRecording() {
			span.End()
		}
		return nil
	}
}

func ApmBeforeTrace() resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		client.SetTransport(apmhttp.WrapRoundTripper(
			http.DefaultClient.Transport,
			apmhttp.WithClientRequestName(func(request *http.Request) string {
				return request.Context().Value(_SpanName).(string)
			}),
			apmhttp.WithClientTrace(),
		))
		return nil
	}
}

func BeforeTrace(tp trace.TracerProvider) resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		client.SetTransport(otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp), otelhttp.WithFilter(
			func(request *http.Request) bool {
				if !parseCtxValOpenTrace(request.Context()) {
					return false
				}
				return true
			})))
		return nil
	}
}

type microNext struct{}

func BeforeMicroSelect(serviceName string, r registry.Registry, options ...selector.SelectOption) resty.RequestMiddleware {
	s := selector.NewSelector(selector.Registry(cache.New(r)), selector.SetStrategy(selector.RoundRobin))
	return func(client *resty.Client, request *resty.Request) error {

		v := request.Context().Value(microNext{})
		if v != nil {
			next := v.(selector.Next)
			node, err := next()
			if err != nil {
				return err
			}
			client.SetHostURL("http://" + node.Address)
			return nil
		}

		next, err := s.Select(serviceName, options...)
		if err != nil {
			return err
		}
		node, err := next()
		if err != nil {
			return err
		}
		client.SetHostURL("http://" + node.Address)
		context.WithValue(request.Context(), microNext{}, next)
		return nil
	}
}

func AfterTraceDebug() resty.ResponseMiddleware {
	return func(client *resty.Client, response *resty.Response) error {
		if !parseCtxValOpenTrace(response.Request.Context()) {
			return nil
		}
		traceInfo := new(TraceInfo)
		traceInfo.Request = SetRequest(response.Request)
		traceInfo.Response = SetResponse(response)
		traceInfo.Info = SetInfo(response.Request.TraceInfo())
		traceRaw, _ := json.Marshal(traceInfo)
		span := trace.SpanFromContext(response.Request.Context())
		span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("info"), semconv.ExceptionMessageKey.String(string(traceRaw))))
		span.End()
		span = trace.SpanFromContext(response.Request.Context())
		return nil
	}
}

//func BeforSelectNode() resty.RequestMiddleware {
//	return func(client *resty.Client, request *resty.Request) error {
//	}
//}
