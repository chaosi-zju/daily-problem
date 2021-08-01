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

	userRouter := r.Group("/api/user")
	userRouter.Use(middleware.JWTAuth())

	userRouter.GET("/overview", handler.GetTodayOverview)
	userRouter.GET("/finish_info", handler.GetFinishInfo)
	userRouter.GET("/problem/daily", handler.GetDailyProblem)
	userRouter.POST("/problem/add", handler.AddProblem)
	userRouter.POST("/problem/update", handler.UpdateProblem)
	userRouter.GET("/problem/finish", handler.FinishProblem)
	userRouter.GET("/problem/unplanned", handler.GetAllUnPlanned)
	userRouter.GET("/problem/get", handler.GetProblemByID)
	userRouter.GET("/problem/pickmore", handler.PickMoreProblem)
	userRouter.GET("/problem/types", handler.GetAllTypes)

	userRouter.GET("/problem/plan/all", handler.GetAllPlanned)
	userRouter.GET("/problem/plan/add", handler.AddToUserPlan)
	userRouter.GET("/problem/plan/remove", handler.RemoveFromPlan)
	userRouter.GET("/problem/plan/doitnow", handler.AddToDaily)

	return r
}
