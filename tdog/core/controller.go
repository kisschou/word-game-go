package core

import (
	"fmt"
)

type Controller struct {
	Req  *Request
	Name string
}

func (c *Controller) SayHi() {
	fmt.Println("You extends core/controller!", c.Name)
}
