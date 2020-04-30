package controller

import (
	"fmt"

	"wordgame/tdog/core"
)

type Member struct {
	Base core.Controller
}

func (member *Member) Login() {
	fmt.Println(member.Base, "\r\n")
	fmt.Println("Login!")
	member.Base.SayHi()
}

func (member *Member) Hello() {
	fmt.Println("Hello World!")
	member.Base.SayHi()
}
