package core

import (
	"fmt"
)

type Controller struct {
}

func (c *Controller) SayHi() {
	fmt.Println("You extends core/controller!")
}
