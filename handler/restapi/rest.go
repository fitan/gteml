package restapi

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/rest"
	"github.com/gin-gonic/gin"
	"sync"
)

type ApiRest struct {
	rest.Restful
}

func NewApiRest(baseRest rest.Restful) *ApiRest {
	return &ApiRest{Restful: baseRest}
}

func (a *ApiRest) Wrap(ctx *gin.Context, data interface{}, err error) {
	res := make(map[string]interface{}, 0)

	if err != nil {
		res["err"] = err.Error()
		res["code"] = 5003
		ctx.JSON(200, res)
		return
	}

	res["data"] = data
	res["code"] = 2000
	ctx.JSON(200, res)
	return
}

type RestfulObj struct {
	Users    *ApiRest
	Roles    *ApiRest
	Services *ApiRest
}

func NewRestfulObj() *RestfulObj {
	db := core.GetCorePool().GetObj().Dao.Storage().DB()
	return &RestfulObj{
		Users:    &ApiRest{rest.NewBaseRest(db, &UserObj{})},
		Roles:    &ApiRest{rest.NewBaseRest(db, &RolesObj{})},
		Services: &ApiRest{rest.NewBaseRest(db, &ServiceObj{})},
	}
}

var restfulAll *RestfulObj
var once sync.Once

func GetRestfulAll() *RestfulObj {
	once.Do(
		func() {
			restfulAll = NewRestfulObj()
		})
	return restfulAll
}
