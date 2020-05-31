package core

import (
	"encoding/json"
	"fmt"

	"wordgame/tdog/lib"
)

type (
	Feign struct {
		Method    string
		BaseUrl   string
		ActionUrl string
		Header    map[string]string
		Body      map[string]interface{}
	}
)

func NewFeign() *Feign {
	feign := &Feign{}
	return feign
}

func (feign *Feign) Decoder(data string) *Feign {
	dataMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		fmt.Println(err)
	}

	feign.Method = dataMap["method"].(string)
	feign.BaseUrl = dataMap["base_url"].(string)
	feign.ActionUrl = dataMap["action_url"].(string)
	headerMap := make(map[string]string)
	json.Unmarshal([]byte(dataMap["header"].(string)), &headerMap)
	feign.Header = headerMap
	bodyMap := make(map[string]interface{})
	json.Unmarshal([]byte(dataMap["body"].(string)), &bodyMap)
	feign.Body = bodyMap

	return feign
}

func (feign *Feign) Target() {
	ConfigLib := new(lib.Config)
	HttpRequestLib := new(lib.HttpRequest)

	// 请求服务不存在
	if !ConfigLib.Get("api_url." + feign.BaseUrl).IsExists() {
	}

	url := ConfigLib.Get("api_url."+feign.BaseUrl).String() + feign.ActionUrl
	HttpRequestLib.Method = feign.Method
	HttpRequestLib.Header = feign.Header
	HttpRequestLib.Url = url
	HttpRequestLib.Params = feign.Body
	res, err := HttpRequestLib.FormRequest()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("===========================> " + HttpRequestLib.Method + " " + HttpRequestLib.Url)
	fmt.Println(res)
}
