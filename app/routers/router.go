package routers

import (
	"wordgame/app/controllers"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var VerificationCodeController controllers.VerificationCode
	verificationCodeRouter := r.Group("/verification_code", &VerificationCodeController.Base)
	{
		verificationCodeRouter.POST("/get", VerificationCodeController.Get)
		verificationCodeRouter.POST("/getPic", VerificationCodeController.GetPic)
		verificationCodeRouter.POST("/verify", VerificationCodeController.Verify)
	}

	return r
}
