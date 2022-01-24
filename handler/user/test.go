package user

import (
	"github.com/fitan/magic/pkg/types"
	"strconv"
)

type CreateIn struct {
	Body struct {
		Hello string `json:"hello"`
	} `json:"body"`
}

// @GenApi /user [post]
func Create(c *types.Core, in *CreateIn) (string, error) {
	//c.Log.Info("这是 create的开始")
	//c.Log.Sync()
	//
	//log := c.CoreLog.TraceLog("nest 嵌套")
	//log.Info("嵌套info： fsfdf fsdfsd ")
	//defer log.Sync()
	//
	//byId, err := c.Storage.User().GetByIds([]int{1, 9, 10, 13})
	//if err != nil {
	//	return nil, err
	//}
	//
	//c.Apis.Baidu().GetSum()
	////data, err := c.Apis.Baidu().GetRoot()
	////if err != nil {
	////	return "", err
	////}
	//
	//_, ok := c.GinX.GinCtx().GetQuery("status")
	//if !ok {
	//	return nil, errors.New("not find query status")
	//}
	//return byId, nil
	//return data.String(), nil
	return "", nil
}

type SayHelloIn struct {
	Query struct {
		Say string `json:"say" form:"say"`
	} `json:"query"`
	CtxKey struct {
		*JwtKey
	} `json:"ctxKey"`
}

type JwtKey struct {
	JwtUserID uint   `ctxkey:"JwtUserIDKey" binding:"required" json:"jwtUserId"`
	TestValue string `json:"testValue"`
}

func (s *SayHelloIn) ServiceID() (serviceID uint) {
	return s.CtxKey.JwtUserID
}

// @GenApi /say [get]
func SayHello(core *types.Core, in *SayHelloIn) (string, error) {

	if in.Query.Say != "" {

		req, err := core.GetApis().Gteml().GetRoot()
		if err != nil {
			return "false", err
		}
		return req.String(), nil
	}

	return core.GetServices().User().Read() + strconv.Itoa(int(in.CtxKey.JwtUserID)) + in.CtxKey.TestValue, nil
}
