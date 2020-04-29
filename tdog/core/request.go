package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	Request struct {
		Header map[string]string
		Params map[string][]string
		Get    map[string][]string
		Post   map[string][]string
		Put    map[string]string
	}
)

func (r *Request) Recv(c *Context) {
	for k, v := range c.Req.Header {
		r.Header[k] = v[0]
	}
	fmt.Println(r.Header)

	for k, v := range c.Req.URL.Query() {
		r.Get[k] = v
	}
	fmt.Println(r.Get)

	if _, ok := c.Req.Header["Content-Type"]; ok {
		if strings.Contains(c.Req.Header["Content-Type"][0], "json") {
			decoder := json.NewDecoder(c.Req.Body)
			var jsonParams map[string]string
			decoder.Decode(&jsonParams)
			r.Put = jsonParams
			fmt.Println(r.Put)
		}

		if strings.Contains(c.Req.Header["Content-Type"][0], "x-www-form-urlencoded") {
			for k, v := range c.Req.PostForm {
				r.Post[k] = v
			}
			fmt.Println(r.Post)
		}

		if strings.Contains(c.Req.Header["Content-Type"][0], "form-data") {
			for k, v := range c.Req.PostForm {
				r.Post[k] = v
			}
			fmt.Println(r.Post)
		}
	}

	fmt.Println(r)
}
