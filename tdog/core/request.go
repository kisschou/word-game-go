package core

import (
	"encoding/json"
	"reflect"
	"strings"
)

type (
	Request struct {
		Header  map[string][]string
		Params  map[string][]string
		Get     map[string][]string
		Post    map[string][]string
		Put     map[string]interface{}
		IsGet   bool
		IsPost  bool
		IsPut   bool
		IsOpt   bool
		IsDel   bool
		IsPatch bool
	}
)

func (r *Request) New(c *Context) {
	var ghostMap map[string][]string
	ghostMap = make(map[string][]string)
	for k, v := range c.Req.Header {
		ghostMap[k] = v
	}
	r.Header = ghostMap

	ghostMap = make(map[string][]string)
	for k, v := range c.Req.URL.Query() {
		ghostMap[k] = v
	}
	r.Get = ghostMap

	if _, ok := c.Req.Header["Content-Type"]; ok {
		if strings.Contains(c.Req.Header["Content-Type"][0], "json") {
			decoder := json.NewDecoder(c.Req.Body)
			var jsonParams map[string]interface{}
			decoder.Decode(&jsonParams)
			r.Put = jsonParams
		}

		if strings.Contains(c.Req.Header["Content-Type"][0], "x-www-form-urlencoded") {
			ghostMap = make(map[string][]string)
			for k, v := range c.Req.PostForm {
				ghostMap[k] = v
			}
			r.Post = ghostMap
		}

		if strings.Contains(c.Req.Header["Content-Type"][0], "form-data") {
			ghostMap = make(map[string][]string)
			for k, v := range c.Req.PostForm {
				ghostMap[k] = v
			}
			r.Post = ghostMap
		}
	}

	// 判断请求类型
	r = checkReqMethod(r, c.Req.Method)

	// 合并参数到Params
	r = merge2Params(r)

	// set to base controller.
	c.BaseController.Req = r
}

func checkReqMethod(r *Request, method string) *Request {
	switch method {
	case "GET":
		r.IsGet = true
		break
	case "POST":
		r.IsPost = true
		break
	case "PUT":
		r.IsPut = true
		break
	case "DELETE":
		r.IsDel = true
		break
	case "OPTIONS":
		r.IsOpt = true
		break
	case "PATCH":
		r.IsPatch = true
		break
	}
	return r
}

func merge2Params(r *Request) *Request {
	params := make(map[string][]string)
	if len(r.Get) > 0 {
		for k, v := range r.Get {
			params[k] = v
		}
	}
	if len(r.Post) > 0 {
		for k, v := range r.Post {
			params[k] = v
		}
	}
	if len(r.Put) > 0 {
		for k, v := range r.Put {
			if reflect.TypeOf(v).Kind().String() == "map" {
				dataJson, _ := json.Marshal(v.(map[string]interface{}))
				params[k] = []string{string(dataJson)}
			} else {
				params[k] = []string{v.(string)}
			}
		}
	}
	r.Params = params
	return r
}
