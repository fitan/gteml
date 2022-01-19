package ginx

import (
	"github.com/fitan/magic/dao/dal/model"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

type Audit struct {
}

func (a *Audit) BindValBefor(core *types.Core) bool {
	return true
}

func (a *Audit) BindValAfter(core *types.Core) bool {
	return true
}

func (a *Audit) BindFnAfter(core *types.Core) bool {
	return true
}

func (a *Audit) Forever(core *types.Core) {
	log := core.GetCoreLog().TraceLog("pkg.ginx.audit")
	var err error
	defer func() {
		if err != nil {
			log.Error("audit inset error", zap.Error(err))
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{}))
		}

		log.Sync()
	}()
	method := core.GinX.GinCtx().Request.Method
	url := core.GinX.GinCtx().Request.RequestURI
	query := core.GinX.GinCtx().Request.URL.Query().Encode()
	remoteIP := core.GinX.GinCtx().Request.RemoteAddr
	requestBuf := make([]byte, 1024)
	n, _ := core.GinX.GinCtx().Request.Body.Read(requestBuf)
	request := requestBuf[0:n]

	responseBuf := make([]byte, 1024)
	n, _ = core.GinX.GinCtx().Request.Response.Body.Read(responseBuf)
	response := responseBuf[0:n]

	statusCode := core.GinX.GinCtx().Request.Response.StatusCode

	err = core.GetServices().Audit().InsetAudit(&model.Audit{
		Url:        url,
		Query:      query,
		Method:     method,
		Request:    string(request),
		Response:   string(response),
		Header:     "",
		StatusCode: statusCode,
		RemoteIP:   remoteIP,
	})

}
