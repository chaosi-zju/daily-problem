package app

import (
	"github.com/chaosi-zju/daily-problem/internal/app/handler"
	"github.com/chaosi-zju/daily-problem/internal/app/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func SetupRoutes() *gin.Engine {
	mode := viper.GetString("gin.mode")
	gin.SetMode(mode)

	r := gin.Default()
	r.Use(middleware.Cors())

	r.Use(static.Serve("/", static.LocalFile("./website", true)))
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")

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
