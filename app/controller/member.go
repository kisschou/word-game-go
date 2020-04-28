package controller

import (
	"fmt"

	"wordgame/tdog/core"
)

type Member struct {
	core.Controller
}

func (member *Member) Login() {
	fmt.Println("Hello World!")
	member.SayHi()
}
