package controller

import (
	"fmt"

	"wordgame/app/service"
	"wordgame/tdog/core"
)

type Member struct {
	Base core.Controller
}

func (member *Member) Login() {
	username := member.Base.Req.Params["username"][0]
	password := member.Base.Req.Params["password"][0]
	fmt.Println("Login with username: [" + username + "] password: [" + password + "]")

	MemberService := new(service.Member)
	MemberService.Login(username, password)
	member.Base.SayHi()
	member.Base.Res.JSON(200, core.H{
		"message": "login success",
	})
}

func (member *Member) Ping() {
	fmt.Println("Hello World!")
	member.Base.SayHi()
	member.Base.Res.String(200, "Pong")
}
