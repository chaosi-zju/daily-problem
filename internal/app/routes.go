package app

import (
	"github.com/chaosi-zju/daily-problem/internal/app/handler"
	"github.com/chaosi-zju/daily-problem/internal/app/middleware"
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

	r.POST("/api/login", handler.Login)
	r.POST("/api/register", handler.Register)
	r.GET("/api/problem/add", handler.AddProblem)
	r.GET("/api/problem/update", handler.UpdateProblem)

	userRouter := r.Group("/api/user")
	userRouter.Use(middleware.JWTAuth())
	userRouter.GET("/problem/daily", handler.GetDailyProblem)
	userRouter.GET("/problem/finish", handler.FinishProblem)
	userRouter.GET("/problem/notredo", handler.ShouldNotRedo)

	return r
}
