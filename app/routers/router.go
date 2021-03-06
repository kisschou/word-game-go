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
		memberRouter.GET("/ping", MemberController.Ping)
	}

	var AuthController controllers.Auth
	authRouter := r.Group("/auth", &AuthController.Base)
	{
		authRouter.POST("getToken", AuthController.GetToken)
	}

	var FeignController controllers.Feign
	feignRouter := r.Group("/feign", &FeignController.Base)
	{
		feignRouter.POST("/http", FeignController.Http)
	}

	var DemoController controllers.Demo
	demoRouter := r.Group("/demo", &DemoController.Base)
	{
		demoRouter.GET("hello", DemoController.Hello)
	}

	return r
}
