package core

import (
	"fmt"
)

type Request struct {
}

func (r *Request) Recv(c *Context) {
	fmt.Println(c.Req)
	fmt.Println(c.Req.URL.Query())
	fmt.Println(c.Writer)
	fmt.Println(c.Params)
}
