package core

import (
	"fmt"

	"wordgame/tdog/lib"
)

type Model struct {
	Sql lib.MySql
}

func (model *Model) SayHi() {
	fmt.Println("You extends core/model!")
}
