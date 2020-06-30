package core

import (
	"fmt"
)

type (
	H map[string]interface{}

	Controller struct {
		Req    *Request
		Res    *Response
		UserId int64
	}
)

func (c *Controller) SayHi() {
	fmt.Println("You extends core/controller!")
}
