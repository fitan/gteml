package httpclient

import (
	"context"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

const _OpenTrace string = "_openTrace"
const _SpanName string = "_spanName"

func NewTraceClient(tracer types.Tracer, client *resty.Client) *TraceClient {
	return &TraceClient{tracer, client}

}

type TraceClient struct {
	types.Tracer
	*resty.Client
}

func (h *TraceClient) SetTracer(t types.Tracer) {
	h.Tracer = t
}

func (r *TraceClient) openTraceCtx(name string) context.Context {
	ctx := context.WithValue(r.Ctx(), _OpenTrace, true)
	ctx = context.WithValue(ctx, _SpanName, name)
	return ctx
}

func (r *TraceClient) offTraceCtx() context.Context {
	return context.WithValue(r.Ctx(), _OpenTrace, false)
}

func (h *TraceClient) R(name string) *resty.Request {
	r := h.Client.R()
	if h.Tracer.IsOpen() {
		r = r.SetContext(h.openTraceCtx(name))
	} else {
		r = r.SetContext(h.offTraceCtx())
	}
	return r
	//return &TraceRequest{
	//	Tracer:  h.Tracer,
	//	Request: r,
	//}
}

//type TraceRequest struct {
//	types.Tracer
//	*resty.Request
//}
//
//func (r *TraceRequest) openTraceCtx(name string) context.Context {
//	ctx := context.WithValue(r.Ctx(), _OpenTrace, true)
//	ctx = context.WithValue(ctx, _SpanName, name)
//	return ctx
//}
//
//func (r *TraceRequest) offTraceCtx() context.Context {
//	return context.WithValue(r.Ctx(), _OpenTrace, false)
//}
//
//func (r *TraceRequest) Get(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Get(url)
//		//return r.Request.SetContext(r.openTraceCtx(name)).Get(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Get(url)
//}
//
//func (r *TraceRequest) Post(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Post(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Post(url)
//}
//
//func (r *TraceRequest) Put(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Put(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Put(url)
//}
//
//func (r *TraceRequest) Delete(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Delete(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Delete(url)
//}
//
//func (r *TraceRequest) Head(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Head(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Head(url)
//}
//
//func (r *TraceRequest) Patch(url string, name string) (*resty.Response, error) {
//	if r.Tracer.IsOpen() {
//		return r.Request.SetContext(r.openTraceCtx(name)).Patch(url)
//	}
//	return r.Request.SetContext(r.offTraceCtx()).Patch(url)
//}
