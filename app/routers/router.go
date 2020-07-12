package routers

import (
	"wordgame/app/controllers"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var LoggerController controllers.Logger
	loggerRouter := r.Group("/logger", &LoggerController.Base)
	{
		loggerRouter.POST("/get", LoggerController.Get)
		loggerRouter.POST("/set", LoggerController.Set)
		loggerRouter.GET("/ping", LoggerController.Ping)
	}

	return r
}
