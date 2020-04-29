package main

import (
	"wordgame/app/controller"
	"wordgame/tdog/core"
)

func main() {
	r := core.NewEngine()

	v1 := r.Group("/v1")
	{
		var member controller.Member
		v1.GET("/login", member.Login)
		v1.POST("/login", member.Login)
		v1.GET("/hello", member.Hello)
	}

	r.Run(":8000")
}
