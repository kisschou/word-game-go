package routers

import (
	"wordgame/app/controllers"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var MemberController controllers.Member
	r.GET("/ping", MemberController.Ping, &MemberController.Base)
	r.POST("/ping", MemberController.Ping, &MemberController.Base)

	memberRouter := r.Group("/member", &MemberController.Base)
	{
		memberRouter.POST("/login", MemberController.Login)
		memberRouter.POST("/info", MemberController.GetInfo)
		memberRouter.POST("/register", MemberController.Register)
		memberRouter.POST("/updateInfo", MemberController.UpdateInfo)
		memberRouter.GET("/ping", MemberController.Ping)
	}

	return r
}
