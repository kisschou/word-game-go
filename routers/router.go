package routers

import (
	"github.com/gin-gonic/gin"
	"wordgame/app/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/public", "./public")
	// r.LoadHTMLGlob("view/**/*")

	r.GET("/ping", controllers.Ping)

	v1 := r.Group("/api")
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/login", controllers.Login)
		v1.POST("/register", controllers.Register)
	}

	return r
}
