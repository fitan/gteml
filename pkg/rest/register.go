package rest

import "github.com/gin-gonic/gin"

func RegisterRestApi(r gin.IRouter, rest Restful, path string) {
	// getlist
	r.GET(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetList)
	})
	// getone
	r.GET(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.GetOne)
	})

	r.POST(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Create)
	})

	r.PUT(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Update)
	})

	r.PUT(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.UpdateMany)
	})

	r.DELETE(path+"/:id", func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.Delete)
	})

	r.DELETE(path, func(ctx *gin.Context) {
		rest.Wrap(ctx, rest.DeleteMany)
	})
}
