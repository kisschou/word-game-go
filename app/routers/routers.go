package routers

import (
	"github.com/gin-gonic/gin"

	"wordgame/app/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	MemberController := new(controller.Member)
	r.GET("/ping", MemberController.Ping)
	memberRouter := r.Group("/member")
	{
		memberRouter.POST("/login", MemberController.Login)
	}

	return r
}
