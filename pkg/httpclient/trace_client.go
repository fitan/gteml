package httpclient

import (
	"context"
	"github.com/fitan/gteml/pkg/common"
	"github.com/go-resty/resty/v2"
)

const _OpenTrace string = "_openTrace"

func NewTraceClient(tracer common.Tracer, client *resty.Client) *TraceClient {
	return &TraceClient{tracer, client}

}

type TraceClient struct {
	common.Tracer
	*resty.Client
}

func (h *TraceClient) SetTracer(t common.Tracer) {
	h.Tracer = t
}

func (h *TraceClient) R() *TraceRequest {
	r := h.Client.R()
	return &TraceRequest{
		Tracer:  h.Tracer,
		Request: r,
	}
}

type TraceRequest struct {
	common.Tracer
	*resty.Request
}

func (r *TraceRequest) openTraceCtx(name string) context.Context {
	return context.WithValue(r.SpanCtx(name), _OpenTrace, true)
}

func (r *TraceRequest) offTraceCtx() context.Context {
	return context.WithValue(r.Request.Context(), _OpenTrace, false)
}

func (r *TraceRequest) Get(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Get(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Get(url)
}

func (r *TraceRequest) Post(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Post(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Post(url)
}

func (r *TraceRequest) Put(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Put(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Put(url)
}

func (r *TraceRequest) Delete(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Delete(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Delete(url)
}

func (r *TraceRequest) Head(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Head(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Head(url)
}

func (r *TraceRequest) Patch(url string, name string) (*resty.Response, error) {
	if r.Tracer.IsOpen() {
		return r.Request.SetContext(r.openTraceCtx(name)).Patch(url)
	}
	return r.Request.SetContext(r.offTraceCtx()).Patch(url)
}
