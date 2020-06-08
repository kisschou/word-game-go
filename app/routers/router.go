package routers

import (
	"wordgame/app/controllers"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var AuthController controllers.Auth
	authRouter := r.Group("/auth", &AuthController.Base)
	{
		// 获取token
		authRouter.GET("token", AuthController.GetToken)
		// 刷新token
		authRouter.GET("refresh", AuthController.RefreshToken)
		// token校验
		authRouter.POST("verify", AuthController.VerifyToken)
		// 登录校验
		authRouter.POST("verifyLogin", AuthController.VerifyLogin)
		// 获取指定参数
		authRouter.POST("getKey", AuthController.GetKey)
		// 获取所有参数
		authRouter.POST("getAllKeys", AuthController.GetAllKeys)
	}

	return r
}
