package main

import (
	"wordgame/app/controller"
	"wordgame/tdog/core"
)

func main() {
	r := core.NewEngine()

	var member controller.Member
	memberRouter := r.Group("/member", &member.Base)
	{
		memberRouter.GET("/login", member.Login)
		memberRouter.POST("/login", member.Login)
		memberRouter.GET("/hello", member.Hello)
	}

	r.Run(":8000")
}
