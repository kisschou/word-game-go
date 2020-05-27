package core

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type (
	Response struct {
		context *Context
	}
)

func (r *Response) New(c *Context) {
	r.context = c
	c.BaseController.Res = r
}

func (r *Response) JSON(code int, obj interface{}) {
	if code >= 0 {
		r.context.Writer.WriteHeader(code)
	}
	r.context.Writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(r.context.Writer)
	if code != http.StatusOK {
		ErrorCore := new(Error)
		newObj := make(map[string]interface{})
		newObj = obj.(H)
		if _, ok := newObj["code"]; ok {
			newObj["message"] = ErrorCore.GetError(newObj["code"].(string))
			delete(newObj, "code")
		}
		obj = newObj
	}
	if err := encoder.Encode(obj); err != nil {
		r.context.Error(err, obj)
		http.Error(r.context.Writer, err.Error(), 500)
	}
}

func (r *Response) XML(code int, obj interface{}) {
	if code >= 0 {
		r.context.Writer.WriteHeader(code)
	}
	r.context.Writer.Header().Set("Content-Type", "application/xml")
	encoder := xml.NewEncoder(r.context.Writer)
	if err := encoder.Encode(obj); err != nil {
		r.context.Error(err, obj)
		http.Error(r.context.Writer, err.Error(), 500)
	}
}

func (r *Response) String(code int, msg string) {
	r.context.Writer.Header().Set("Content-Type", "text/plain")
	r.context.Writer.WriteHeader(code)
	if code != http.StatusOK {
		ErrorCore := new(Error)
		msg = ErrorCore.GetError(msg)
	}
	r.context.Writer.Write([]byte(msg))
}

func (r *Response) Data(code int, data []byte) {
	r.context.Writer.WriteHeader(code)
	r.context.Writer.Write(data)
}
