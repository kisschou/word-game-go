package core

import (
	"fmt"
)

type (
	H map[string]interface{}

	Controller struct {
		Req  *Request
		Res  *Response
		Name string
	}
)

func (c *Controller) SayHi() {
	fmt.Println("You extends core/controller!", c.Name)
}
