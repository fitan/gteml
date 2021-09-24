package httpclient

import (
	"github.com/go-resty/resty/v2"
	"net"
	"time"
)

type TraceInfo struct {
	Response *Response `json:"response,omitempty"`

	Request *Request `json:"request,omitempty"`

	Info *Info `json:"trace_info,omitempty"`
}

type Request struct {
	Url        string      `json:"url"`
	Method     string      `json:"method"`
	QueryParam string      `json:"query_param"`
	Body       interface{} `json:"body"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
	Proto      string      `json:"proto"`
	ReceivedAt string      `json:"received_at"`
}

type Info struct {
	DNSLookup      time.Duration `json:"dns_lookup"`
	ConnTime       time.Duration `json:"conn_time"`
	TCPConnTime    time.Duration `json:"tcp_conn_time"`
	TLSHandshake   time.Duration `json:"tls_handshake"`
	ServerTime     time.Duration `json:"server_time"`
	ResponseTime   time.Duration `json:"response_time"`
	TotalTime      time.Duration `json:"total_time"`
	IsConnReused   bool          `json:"is_conn_reused"`
	IsConnWasIdle  bool          `json:"is_conn_was_idle"`
	ConnIdleTime   time.Duration `json:"conn_idle_time"`
	RequestAttempt int           `json:"request_attempt"`
	RemoteAddr     net.Addr      `json:"remote_addr"`
}

func SetResponse(resp *resty.Response) *Response {
	data := new(Response)
	data.StatusCode = resp.StatusCode()
	data.Proto = resp.Proto()
	data.ReceivedAt = resp.ReceivedAt().String()
	data.Body = resp.Body()
	return data
}

func SetRequest(req *resty.Request) *Request {
	data := new(Request)
	data.Body = req.Body
	data.Url = req.URL
	data.Method = req.Method
	data.QueryParam = req.QueryParam.Encode()
	return data
}

func SetInfo(traceInfo resty.TraceInfo) *Info {
	return &Info{
		DNSLookup:      traceInfo.DNSLookup,
		ConnTime:       traceInfo.ConnTime,
		TCPConnTime:    traceInfo.TCPConnTime,
		TLSHandshake:   traceInfo.TLSHandshake,
		ServerTime:     traceInfo.ServerTime,
		ResponseTime:   traceInfo.ResponseTime,
		TotalTime:      traceInfo.TotalTime,
		IsConnReused:   traceInfo.IsConnReused,
		IsConnWasIdle:  traceInfo.IsConnWasIdle,
		ConnIdleTime:   traceInfo.ConnIdleTime,
		RequestAttempt: traceInfo.RequestAttempt,
		RemoteAddr:     traceInfo.RemoteAddr,
	}
}
