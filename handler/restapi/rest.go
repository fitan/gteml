package restapi

import (
	"github.com/fitan/magic/pkg/rest"
	"github.com/gin-gonic/gin"
)

type ApiRest struct {
	rest.Restful
}

func NewApiRest(baseRest rest.Restful) *ApiRest {
	return &ApiRest{Restful: baseRest}
}

func (a *ApiRest) Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error)) {
	res := make(map[string]interface{}, 0)

	data, err := fn(ctx)
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
