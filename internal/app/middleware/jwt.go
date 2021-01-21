package middleware

import (
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/util"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-jwt-token")
		if token == "" {
			util.ResponseHTTPError(c, 401, "no token for authentication")
			c.Abort()
			return
		}

		j := model.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			util.ResponseHTTPError(c, 401, err.Error())
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
