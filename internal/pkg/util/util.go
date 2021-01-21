package util

import (
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/model"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}

func ResponseHTTPError(c *gin.Context, httpcode int, message string) {
	c.JSON(httpcode, gin.H{
		"code":    httpcode,
		"message": message,
		"data":    nil,
	})
}

func GetUserIdFromContext(c *gin.Context) (uint, error) {
	if v, ok := c.Get("claims"); ok {
		if claims, ok := v.(*model.CustomClaims); ok {
			return claims.UserID, nil
		}
	}
	return 0, fmt.Errorf("user info invalid")
}

func GetUserFromContext(c *gin.Context) (model.User, error) {
	var user model.User
	if v, ok := c.Get("claims"); ok {
		if claims, ok := v.(*model.CustomClaims); ok {
			if err := mysqld.Db.First(&user, claims.UserID).Error; err == nil {
				return user, nil
			}
		}
	}
	return user, fmt.Errorf("user info invalid")
}
