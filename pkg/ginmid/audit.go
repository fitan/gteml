package ginmid

import (
	"bytes"
	"fmt"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func WithAuditFilter(fc func(ctx *gin.Context) bool) func(a *Audit) {
	return func(a *Audit) {
		a.filter = fc
	}
}

type Audit struct {
	filter func(ctx *gin.Context) bool
}

func NewAudit(options ...func(a *Audit)) *Audit {
	a := &Audit{
		filter: func(ctx *gin.Context) bool {
			if ctx.Request.Method == "GET" {
				return false
			}
			return true
		},
	}
	for _, o := range options {
		o(a)
	}
	return a
}

func (a *Audit) Audit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if a.filter != nil {
			if !a.filter(ctx) {
				ctx.Next()
				return
			}
		}

		startTime := time.Now()

		var (
			core *types.Core
			err  error
		)

		//defer func() {
		//	log := core.GetCoreLog().ApmLog("pkg.ginmid.audit")
		//	if err != nil {
		//		log.Error(err.Error())
		//	}
		//	log.Sync()
		//}()

		if value, ok := ctx.Get(types.CoreKey); ok {
			core = value.(*types.Core)
		} else {
			err = errors.New("gin ctx not found type: types.Core")
			ctx.JSON(http.StatusInternalServerError, ginx.GinXResult{
				Code: 5000,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		bodyB, _ := ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyB))

		ctx.Next()

		log := core.GetCoreLog().ApmLog("pkg.ginmid.audit")
		defer func() {
			if err != nil {
				log.Error(err.Error())
			}

			log.Sync()
		}()

		method := ctx.Request.Method
		url := ctx.FullPath()
		query := ctx.Request.URL.Query().Encode()
		remoteIP := ctx.Request.RemoteAddr
		response := blw.body.String()
		log.Info(response)
		statusCode := ctx.Writer.Status()
		err = core.GetDao().Storage().Query().WrapQuery().Audit.Create(&model.Audit{
			Url:        url,
			Query:      query,
			Method:     method,
			Request:    string(bodyB),
			Response:   response,
			Header:     "",
			StatusCode: statusCode,
			RemoteIP:   remoteIP,
			ClientIP:   ctx.ClientIP(),
			CostTime:   fmt.Sprintf("%v", time.Now().Sub(startTime)),
		})
	}
}
