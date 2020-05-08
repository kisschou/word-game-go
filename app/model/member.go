package model

import (
	"fmt"
	"wordgame/tdog/core"
)

type Member struct {
	core.Model
}

func (member *Member) Login() {
	fmt.Println("Hello World!")
	member.SayHi()
}
