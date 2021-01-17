package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
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
