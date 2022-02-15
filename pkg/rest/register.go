package rest

import "github.com/gin-gonic/gin"

func RegisterRestApi(r gin.IRouter, rest Restful, path string) {
	// /path?ids=1&ids2&_page=1&_limit=2&_sort=name&_order=dec&filter={"name":"bo"}
	r.GET(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetList)
	})
	// /path/1
	r.GET(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetOne)
	})

	// /path
	r.POST(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Create)
	})

	// /path/1
	r.PUT(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Update)
	})

	// /path?ids=1?ids=2
	r.PUT(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.UpdateMany)
	})

	// /path/1
	r.DELETE(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Delete)
	})

	// /path?ids=1?ids=2
	r.DELETE(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.DeleteMany)
	})

	// 查询某个字段 一般用作 ui input 的 selecte 或者 autoComplete
	// /path/fields/name?_keyWord=an
	r.GET(path+"/fields/:name", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetField)
	})

	// /path/fields?_keyWord=an?fields=xx?fields=xx
	r.GET(path+"/fields", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetFields)
	})

	// /path/1/relateion/roles?_page
	r.GET("path"+"/:id/relations/:relationName", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Relations)
	})
}
