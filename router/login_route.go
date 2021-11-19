package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func LoginRoute(r gin.IRouter, jwtMid *jwt.GinJWTMiddleware) {
	r.POST("/login", jwtMid.LoginHandler)
	r.POST("logout", jwtMid.LogoutHandler)
	r.GET("/refresh_token", jwtMid.RefreshHandler)
}
