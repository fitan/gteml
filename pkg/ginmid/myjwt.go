package ginmid

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/fitan/magic/dao/dal/model"
	core2 "github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var identityKey = "id"

type loginValues struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	core := core2.GetCorePool().GetObj()
	jwtConf := core.GetConfig().Jwt
	realm := jwtConf.Realm
	key := jwtConf.SecretKey

	timeout, err := time.ParseDuration(jwtConf.Timeout)
	if err != nil {
		log.Panicln(err)
	}
	maxRefresh, err := time.ParseDuration(jwtConf.MaxRefresh)
	if err != nil {
		log.Panicln(err)
	}

	tokenHeadName := jwtConf.TokenHeadName

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      realm,
		Key:        []byte(key),
		Timeout:    timeout,
		MaxRefresh: maxRefresh,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  tokenHeadName + " " + token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			login := loginValues{}
			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			core := c.MustGet(types.CoreKey).(types.ServiceCore)

			user, err := core.GetServices().User().Login(login.UserName, login.Password)
			return user, err

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			c.Set(types.JwtUserIDKey, uint(data.(float64)))
			return true
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{types.JwtUserIDKey: v.ID}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userID := claims[types.JwtUserIDKey].(float64)
			return userID
		},
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: tokenHeadName,
		TimeFunc:      time.Now,
	})
}
