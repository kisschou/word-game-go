package core

import (
	"encoding/json"
	"reflect"

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
		LoggerLib := new(lib.Logger)
		LoggerLib.Level = 0
		LoggerLib.Key = "error"
		LoggerLib.New(err.Error())
	}

	CryptLib := new(lib.Crypt)
	UtilLib := new(lib.Util)
	feign.Method = ""
	if UtilLib.InMap("method", dataMap) {
		feign.Method = dataMap["method"].(string)
	}
	feign.BaseUrl = ""
	if UtilLib.InMap("base_url", dataMap) {
		feign.BaseUrl = dataMap["base_url"].(string)
	}
	feign.ActionUrl = ""
	if UtilLib.InMap("action_url", dataMap) {
		feign.ActionUrl = dataMap["action_url"].(string)
	}
	headerMap := make(map[string]string)
	if UtilLib.InMap("header", dataMap) {
		if reflect.TypeOf(dataMap["header"]).Kind().String() == "map" {
			for k, v := range dataMap["header"].(map[string]interface{}) {
				headerMap[k] = v.(string)
			}
		} else {
			CryptLib.Str = dataMap["header"].(string)
			json.Unmarshal([]byte(CryptLib.UrlBase64Decode()), &headerMap)
		}
	}
	feign.Header = headerMap
	bodyMap := make(map[string]interface{})
	if UtilLib.InMap("body", dataMap) {
		if reflect.TypeOf(dataMap["body"]).Kind().String() == "map" {
			bodyMap = dataMap["body"].(map[string]interface{})
		} else {
			CryptLib.Str = dataMap["body"].(string)
			json.Unmarshal([]byte(CryptLib.UrlBase64Decode()), &bodyMap)
		}
	}
	feign.Body = bodyMap

	return feign
}

func (feign *Feign) Target() (code int, res string) {
	ConfigLib := new(lib.Config)
	HttpRequestLib := new(lib.HttpRequest)

	// 请求服务不存在
	if !ConfigLib.Get("api_url." + feign.BaseUrl).IsExists() {
		code = 0
		res = "ERROR_FEIGN_SERVICE_MISSING"
		return
	}

	url := ConfigLib.Get("api_url."+feign.BaseUrl).String() + feign.ActionUrl
	HttpRequestLib.Method = feign.Method
	HttpRequestLib.Header = feign.Header
	HttpRequestLib.Url = url
	HttpRequestLib.Params = feign.Body
	code, res, err := HttpRequestLib.FormRequest()
	if err != nil {
		code = 0
		res = "ERROR_FEIGN_REQUEST_FAIL"

		LoggerLib := new(lib.Logger)
		LoggerLib.Level = 0
		LoggerLib.Key = "error"
		LoggerLib.New(err.Error())
		return
	}
	return
}
