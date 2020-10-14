package routers

import (
	"wordgame/app/controllers"

	"github.com/kisschou/tdog"
)

func InitRouter() *tdog.HttpEngine {
	r := tdog.NewEngine()

	var AuthController controllers.Auth
	authRouter := r.Group("/auth", &AuthController.Base)
	{
		authRouter.GET("getToken", AuthController.GetToken)
	}

	var FeignController controllers.Feign
	feignRouter := r.Group("/feign", &FeignController.Base)
	{
		feignRouter.POST("/http", FeignController.Http)
		feignRouter.GET("/ping", FeignController.Ping)
	}

	return r
}
