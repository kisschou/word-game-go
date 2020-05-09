package model

import (
	"fmt"

	"wordgame/tdog/core"
)

type Member struct {
	core.Model
}

func (member *Member) Login(username string, password string) {
	fmt.Println("Hello World!")
	member.SayHi()
}
