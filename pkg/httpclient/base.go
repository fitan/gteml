package httpclient

import (
	"github.com/go-resty/resty/v2"
	"go-micro.dev/v4/registry"
	"go.opentelemetry.io/otel/sdk/trace"
	"net"
	"net/http"
	"time"
)

type option struct {
	Host string

	Tp *trace.TracerProvider
	// 记录所有的详细的http info, 否则只记录发生错误时的http info
	TraceDebug bool

	MicroServiceName string
	MicroRegistry    registry.Registry

	Debug            bool
	TimeOut          time.Duration
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

type Option func(*option)

func NewClient(fs ...Option) *resty.Client {
	o := option{
		Debug:            false,
		TimeOut:          30 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    10 * time.Second,
		RetryMaxWaitTime: 30 * time.Second,
	}
	for _, f := range fs {
		f(&o)
	}

	client := resty.New().SetDebug(o.Debug).SetTimeout(o.TimeOut).SetRetryCount(o.RetryCount).SetRetryWaitTime(o.RetryWaitTime).SetRetryMaxWaitTime(o.RetryMaxWaitTime)

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   10,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client.SetTransport(transport)
	if o.Tp != nil {
		client = client.EnableTrace()
		client = client.OnBeforeRequest(BeforeTrace(o.Tp))
		//client = client.OnBeforeRequest(ApmBeforeTrace())
		if o.TraceDebug {
			client = client.OnAfterResponse(AfterTraceDebug())
		} else {
			client = client.OnError(AfterErrorTrace())
			client = client.OnAfterResponse(AfterErrorSpanEnd())
		}
	}

	if o.Host != "" {
		client = client.SetHostURL(o.Host)
	}

	if o.MicroServiceName != "" {
		client = client.OnBeforeRequest(BeforeMicroSelect(o.MicroServiceName, o.MicroRegistry))
	}

	return client
}

func WithTrace(tp *trace.TracerProvider, traceDebug bool) func(o *option) {
	return func(o *option) {
		o.Tp = tp
		o.TraceDebug = traceDebug
	}
}

func WithHost(host string) Option {
	return func(o *option) {
		o.Host = host
	}
}

func WithMicroHost(serviceName string, r registry.Registry) Option {
	return func(o *option) {
		o.MicroServiceName = serviceName
		o.MicroRegistry = r
	}
}

func WithDebug(debug bool) Option {
	return func(o *option) {
		o.Debug = debug
	}
}

func WithTimeOut(timeOut time.Duration) Option {
	return func(o *option) {
		o.TimeOut = timeOut
	}
}

func WithRetry(retryCount int, retryWaitTime, retryMaxWaitTime time.Duration) Option {
	return func(o *option) {
		o.RetryCount = retryCount
		o.RetryWaitTime = retryWaitTime
		o.RetryMaxWaitTime = retryMaxWaitTime
	}
}
