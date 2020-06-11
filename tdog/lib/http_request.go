package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	HttpRequest struct {
		Method string
		Header map[string]string
		Url    string
		Params map[string]interface{}
	}
)

/** 使用说明
 * header := make(map[string]string)
 * params := make(map[string]interface{})
 * // header
 * header["Authorization"] = resMap["authorization"].(string)
 * header["Content-Type"] = "application/json"
 * header["Connection"] = "keep-alive"
 * // params
 * params["username"] = "admin"
 * params["password"] = "$2a$10$fP.426qKaTmix50Oln8L.uav55gELhAd0Eg66Av4oG86u8km7D/Ky"
 * HttpRequestLib := new(lib.HttpRequest)
 * HttpRequestLib.Method = "POST"
 * HttpRequestLib.Header = header
 * HttpRequestLib.Url = "http://127.0.0.1:8000/member/login"
 * HttpRequestLib.Params = params
 * res, err := HttpRequestLib.FormRequest()
 * if err != nil {
 *    fmt.Println(err)
 * }
 * fmt.Println("===========================> " + HttpRequestLib.Method + " " + HttpRequestLib.Url)
 * fmt.Println(res)
 */
func (hp *HttpRequest) FormRequest() (httpCode int, resData string, err error) {
	client := &http.Client{}

	reqDataJson, _ := json.Marshal(hp.Params)

	req, err := http.NewRequest(hp.Method, hp.Url, bytes.NewBuffer(reqDataJson))
	if err != nil {
		httpCode = http.StatusInternalServerError
		return
	}

	for k, v := range hp.Header {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		httpCode = http.StatusInternalServerError
		return
	}

	httpCode = res.StatusCode
	resData = string(body)
	return
}

/*
 * // 发送二进制数据流
 * // header
 * header["Authorization"] = resMap["authorization"].(string)
 * header["Content-Type"] = "text/plain"
 * header["Connection"] = "Keep-Alive"
 * // params
 * params["method"] = "POST"
 * params["base_url"] = "user_url"
 * params["action_url"] = "/member/login"
 * params["header"] = ""
 * params["body"] = ""
 * sendHeader := make(map[string]string)
 * sendHeader["Authorization"] = resMap["authorization"].(string)
 * sendHeaderJson, _ := json.Marshal(sendHeader)
 * params["header"] = sendHeaderJson
 * sendBody := make(map[string]interface{})
 * sendBody["username"] = "admin"
 * sendBody["password"] = "$2a$10$fP.426qKaTmix50Oln8L.uav55gELhAd0Eg66Av4oG86u8km7D/Ky"
 * sendBodyJson, _ := json.Marshal(sendBody)
 * params["body"] = sendBodyJson
 *
 * HttpRequestLib.Method = "POST"
 * HttpRequestLib.Header = header
 * HttpRequestLib.Url = "http://127.0.0.1:8000/feign/http"
 * HttpRequestLib.Params = params
 * res, err = HttpRequestLib.BytesPost()
 * if err != nil {
 *	fmt.Println(err)
 * }
 * mt.Println("===========================> " + HttpRequestLib.Method + " " + HttpRequestLib.Url)
 * fmt.Println(res)
 */
func (hp *HttpRequest) BytesPost() (int, string, error) {
	data, _ := json.Marshal(hp.Params)
	body := bytes.NewReader(data)
	req, err := http.NewRequest(hp.Method, hp.Url, body)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return http.StatusInternalServerError, "", err
	}

	for k, v := range hp.Header {
		req.Header.Set(k, v)
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return http.StatusInternalServerError, "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
	}

	return resp.StatusCode, string(b), err
}

// 针对请求网关服务构建
// 继承后只需要set结构体中的Params
func (hp *HttpRequest) ServicePost() (bool, string) {
	ConfigLib := new(Config)
	// header
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Connection"] = "Keep-Alive"

	hp.Method = "POST"
	hp.Header = header
	hp.Url = ConfigLib.Get("api_url.gateway_url").String() + "/feign/http"
	httpCode, res, err := hp.BytesPost()
	if httpCode != http.StatusOK || err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return false, ""
	}
	return true, res
}
