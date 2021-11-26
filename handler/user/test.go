package user

import (
	"errors"
	"github.com/fitan/magic/ent"
	"github.com/fitan/magic/pkg/types"
)

type CreateIn struct {
	Body struct {
		Hello string `json:"hello"`
	} `json:"body"`
	Uri    struct{}
	Header struct{}
}

// @GenApi /user [post]
func Create(c *types.Core, in *CreateIn) ([]*ent.User, error) {
	c.Log.Info("这是 create的开始")
	c.Log.Sync()

	log := c.CoreLog.TraceLog("nest 嵌套")
	log.Info("嵌套info： fsfdf fsdfsd ")
	defer log.Sync()

	byId, err := c.Storage.User().GetByIds([]int{1, 9, 10, 13})
	if err != nil {
		return nil, err
	}

	c.Apis.Baidu().GetSum()
	//data, err := c.Apis.Baidu().GetRoot()
	//if err != nil {
	//	return "", err
	//}

	_, ok := c.GinX.GinCtx().GetQuery("status")
	if !ok {
		return nil, errors.New("not find query status")
	}
	return byId, nil
	//return data.String(), nil
}

type SayHelloIn struct {
	Query struct {
		Say string `json:"say" form:"say"`
	} `json:"query"`
}

func (s *SayHelloIn) ServiceID() (serviceID uint) {
	return 1
}

// @GenApi /say [get]
func SayHello(core *types.Core, in *SayHelloIn) (string, error) {

	if in.Query.Say != "" {
		return in.Query.Say, nil
	}

	return core.GetServices().User().Read(), nil
}
