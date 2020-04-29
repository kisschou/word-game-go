package controller

import (
	"fmt"

	"wordgame/tdog/core"
)

type Member struct {
	core.Controller
}

func (member *Member) Login() {
	fmt.Println("Login!")
	member.SayHi()
}

func (member *Member) Hello() {
	fmt.Println("Hello World!")
	member.SayHi()
}
