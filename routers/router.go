package routers

import (
	"github.com/gin-gonic/gin"
	"wordgame/app/controllers"
	"wordgame/library/logger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logger.LoggerAccess())

	r.Static("/public", "./public")
	// r.LoadHTMLGlob("view/**/*")

	r.GET("/ping", controllers.Ping)

	v1 := r.Group("/api")
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/uauth", controllers.UAuth)
		v1.POST("/login", controllers.Login)
		v1.POST("/info", controllers.GetInfo)
		v1.POST("/register", controllers.Register)
	}

	return r
}
