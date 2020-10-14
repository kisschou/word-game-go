package core

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"wordgame/tdog/lib"
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

	// 跨域
	r.context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	r.context.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

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

func (r *Response) Html(data interface{}) {
	path, _ := os.Getwd()
	t, _ := template.ParseFiles(path + "/app/views/" + r.context.Req.RequestURI + ".tpl")
	t.Execute(r.context.Writer, data)
}

func (r *Response) Captcha(code string) {
	d := make([]int, 4)
	for i := 0; i < len(code); i++ {
		iInt, _ := strconv.Atoi(code[i : i+1])
		d[i] = iInt
	}
	r.context.Writer.Header().Set("Content-Type", "image/png")
	lib.NewImage(d, 100, 40).WriteTo(r.context.Writer)
}
