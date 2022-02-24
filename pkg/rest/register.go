package rest

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
)

func getContext(ctx *gin.Context) types.ServiceCore {
	v, _ := ctx.MustGet(types.CoreKey).(types.ServiceCore)
	return v
}

func RegisterRestApi(r gin.IRouter, rest Restful, path string) {
	// /path?ids=1&ids2&_page=1&_limit=2&_sort=name&_order=dec&filter={"name":"bo"}
	r.GET(path, func(ctx *gin.Context) {
		var count int64
		res, err := rest.GetList(getContext(ctx), nil, &count)
		rest.Wrap(ctx, GetListRes{
			Count: count,
			List:  res,
		}, err)
	})
	// /path/1
	r.GET(path+"/:id", func(ctx *gin.Context) {
		res, err := rest.GetOne(getContext(ctx), nil)
		rest.Wrap(ctx, res, err)
	})

	// /path
	r.POST(path, func(ctx *gin.Context) {
		res, err := rest.Create(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// /path/1
	r.PUT(path+"/:id", func(ctx *gin.Context) {
		res, err := rest.Update(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	//// /path?ids=1?ids=2
	//r.PUT(path, func(ctx *gin.Context) {
	//	rest.Wrap(ctx, rest.UpdateMany)
	//})

	// /path/1
	r.DELETE(path+"/:id", func(ctx *gin.Context) {
		res, err := rest.Delete(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// /path?ids=1?ids=2
	r.DELETE(path, func(ctx *gin.Context) {
		res, err := rest.DeleteMany(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// 查询某个字段 一般用作 ui input 的 selecte 或者 autoComplete
	// /path/fields/name?_keyWord=an
	r.GET(path+"/fields/:name", func(ctx *gin.Context) {
		res, err := rest.GetField(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// /path/fields?_keyWord=an?fields=xx?fields=xx
	r.GET(path+"/fields", func(ctx *gin.Context) {
		res, err := rest.GetFields(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// /path/1/relateions/roles?_page
	r.GET(path+"/:id/relations/:relationName", func(ctx *gin.Context) {
		var count int64
		res, err := rest.RelationGet(getContext(ctx), "", nil, &count)
		rest.Wrap(ctx, GetListRes{
			Count: count,
			List:  res,
		}, err)
	})

	r.POST(path+"/:id/relations/:relationName", func(ctx *gin.Context) {
		res, err := rest.RelationCreate(getContext(ctx), "")
		rest.Wrap(ctx, res, err)
	})

	r.PUT(path+"/:id/relations/:relationName", func(ctx *gin.Context) {
		res, err := rest.RelationUpdate(getContext(ctx), "")
		rest.Wrap(ctx, res, err)
	})

	// /path/1/relations?_fields=xx
	r.POST(path+"/:id/relations", func(ctx *gin.Context) {
		res, err := rest.RelationsCreate(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})

	// /path/1/relations?_fields=xx
	r.PUT(path+"/:id/relations", func(ctx *gin.Context) {
		res, err := rest.RelationsUpdate(getContext(ctx))
		rest.Wrap(ctx, res, err)
	})
}
