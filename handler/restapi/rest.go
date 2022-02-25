package restapi

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/restcommon"
	"github.com/gin-gonic/gin"
	"sync"
)

type ApiRest struct {
	restcommon.Restful
}

func NewApiRest(baseRest restcommon.Restful) *ApiRest {
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
	db := core.GetCorePool().GetObj().Dao.DB()
	return &RestfulObj{
		Users:    &ApiRest{restcommon.NewBaseRest(db, &UserObj{})},
		Roles:    &ApiRest{restcommon.NewBaseRest(db, &RolesObj{})},
		Services: &ApiRest{restcommon.NewBaseRest(db, &ServiceObj{})},
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
