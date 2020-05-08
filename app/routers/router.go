package routers

import (
	"wordgame/app/controller"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var member controller.Member
	r.GET("/ping", member.Ping, &member.Base)
	r.POST("/ping", member.Ping, &member.Base)

	memberRouter := r.Group("/member", &member.Base)
	{
		memberRouter.GET("/login", member.Login)
		memberRouter.POST("/login", member.Login)
		memberRouter.GET("/ping", member.Ping)
	}

	return r
}
