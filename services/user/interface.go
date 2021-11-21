package user

import (
	"context"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

func NewUser(core types.ServiceCore) types.Userer {
	return &User{core}
}

type User struct {
	Core types.ServiceCore
}

func (u *User) Login(username string, password string) (*model.User, error) {
	return u.Core.GetDao().Storage().User().CheckPassword(context.Background(), username, password)
}

func (u *User) Read() string {
	//log := u.Core.GetCoreLog().TraceLog("user.read")
	//defer log.Sync()
	//
	//log := u.Core.GetCoreLog().ApmLog("user.read")
	req, _ := u.Core.GetDao().Storage().User().ById(u.Core.GetTrace().SpanCtx("byid"), 1)

	log := u.Core.GetCoreLog().ApmLog("read")
	log.Error("this is read", zap.String("read", "read"), zap.Any("carry", map[string]interface{}{"method1": "1", "method2": "2"}))
	log.Sync()

	log = u.Core.GetCoreLog().ApmLog("read1")
	log.Error("this is read1", zap.String("read1", "read1"), zap.Any("carry", map[string]interface{}{"method1": "1", "method2": "2"}))
	log.Sync()

	r, _ := u.Core.GetApis().Baidu().GetRoot()
	return r.String()

	//log.Error("this is read", zap.String("ceshi", "ceshi"), zap.Any("OutPut", req), zap.Any("InPut", map[string]interface{}{"username": "email"}))
	//time.Sleep(1*time.Second)
	//log.Error("this is read 2")
	//time.Sleep(1*time.Second)
	//log.Error("this is read 3")
	//time.Sleep(1*time.Second)
	//
	//time.Sleep(2*time.Second)
	//defer log.Sync()
	//
	//ctx := u.Core.GetTrace().Ctx()
	//span, ctx := apm.StartSpan(ctx, "service", "service")
	//span.End()
	//span, ctx = apm.StartSpan(ctx, "service1", "service1")
	//span.End()
	//
	//e := apm.DefaultTracer.NewErrorLog(apm.ErrorLogRecord{
	//	Message:       "hello",
	//	MessageFormat: "",
	//	Level:         "error",
	//	LoggerName:    "zap",
	//	Error:         errors.New("errorlogrealod"),
	//})
	//e.SetSpan(span)
	//e.Send()

	//log.Error("this is end",zap.Any("OutPut", req))

	//transaction := apm.DefaultTracer.StartTransaction("GET /", "request")
	//transaction.Result = "Success"
	//transaction.Context.SetLabel("region", "us-east-1")
	//transaction.End()
	//
	//transaction.TraceContext()
	//opts := apm.TransactionOptions{Start: time.Now(), TraceContext: transaction.TraceContext()}
	//t1 := apm.DefaultTracer.StartTransactionOptions("GET /1", "method", opts)
	//t1.End()
	//
	//
	//
	//span := t1.StartSpan("service", "service.user", nil)
	//span.End()
	//
	//
	//spanopts := apm.SpanOptions{Start: time.Now(), Parent: span.TraceContext()}
	//span = apm.DefaultTracer.StartSpan("opts span", "opts span", span.ParentID(), spanopts)
	//span.End()
	//
	//spanopts = apm.SpanOptions{Start: time.Now(), Parent: span.TraceContext()}
	//span = apm.DefaultTracer.StartSpan("error", "error", span.ParentID(), spanopts)
	//
	//
	//e := apm.DefaultTracer.NewErrorLog(apm.ErrorLogRecord{
	//	Message:       "hello",
	//	MessageFormat: "",
	//	Level:         "error",
	//	LoggerName:    "zap",
	//	Error:         errors.New("errorlogrealod"),
	//})
	//e.SetTransaction(transaction)
	//e.Send()

	//var logger = zap.NewExample(zap.WrapCore((&apmzap.Core{}).WrapCore))
	//traceContextFields := apmzap.TraceContext(u.Core.GetTrace().Ctx())
	//logger.With(traceContextFields...).Error("this first user")
	//logger.With(traceContextFields...).Error("this 1 user")

	return req.Email
}

func (u *User) Create() {
	panic("implement me")
}

func (u *User) Update() {
	panic("implement me")
}

func (u *User) Delete() {
	panic("implement me")
}
