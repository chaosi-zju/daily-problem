package app

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pbnjay/memory"
	"github.com/spf13/viper"
	"runtime"
	"strconv"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func SetupRoutes() *gin.Engine {
	mode := viper.GetString("gin.mode")
	gin.SetMode(mode)

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./website", true)))
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")

	r.GET("/api/specs", func(c *gin.Context) {
		c.JSON(200, []string{
			runtime.GOOS,
			strconv.Itoa(runtime.NumCPU()),
			strconv.FormatUint(memory.TotalMemory()/(1024*1024), 10),
		})
	})

	return r
}
