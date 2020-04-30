package routers

import (
	"wordgame/app/controller"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var member controller.Member
	r.POST("/hello", member.Login, &member.Base)

	memberRouter := r.Group("/member", &member.Base)
	{
		memberRouter.GET("/login", member.Login)
		memberRouter.POST("/login", member.Login)
		memberRouter.GET("/hello", member.Hello)
	}

	return r
}
