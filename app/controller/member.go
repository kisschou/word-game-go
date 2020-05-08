package controller

import (
	"fmt"

	"wordgame/tdog/core"
)

type Member struct {
	Base core.Controller
}

func (member *Member) Login() {
	username := member.Base.Req.Params["username"][0]
	password := member.Base.Req.Params["password"][0]
	fmt.Println("Login with username: [" + username + "] password: [" + password + "]")
	member.Base.SayHi()
}

func (member *Member) Hello() {
	fmt.Println("Hello World!")
	member.Base.SayHi()
}
