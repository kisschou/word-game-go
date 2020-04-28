package core

import (
	"fmt"
)

type Model struct {
}

func (model *Model) SayHi() {
	fmt.Println("You extends core/model!")
}
