package controllers

import (
	"encoding/json"
	"net/http"

	"wordgame/app/services"
	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type Member struct {
	Base core.Controller
}

// swagger:operation POST /member/login member login
// ---
// summary: 用户登录
// description: 用户登录
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: username
//   in: body
//   description: 用户名
//   type: string
//   required: true
// - name: password
//   in: body
//   description: 密码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (member *Member) Login() {
	username := ""
	password := ""

	if _, ok := member.Base.Req.Params["username"]; ok {
		if len(member.Base.Req.Params["username"]) > 0 {
			username = member.Base.Req.Params["username"][0]
		}
	}
	if _, ok := member.Base.Req.Params["password"]; ok {
		if len(member.Base.Req.Params["password"]) > 0 {
			password = member.Base.Req.Params["password"][0]
		}
	}

	if len(username) < 5 || len(password) < 6 {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	MemberService := new(services.Member)
	MemberService.IpAddr = member.Base.Req.IpAddr
	memberInfo, err := MemberService.Login(username, password)
	if err != nil {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	authorization := ""
	HttpRequestLib := new(lib.HttpRequest)

	// params
	params := make(map[string]interface{})
	params["method"] = "GET"
	params["base_url"] = "auth_url"
	params["action_url"] = "/auth/token"
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Connection"] = "keep-alive"
	params["header"] = header
	body := make(map[string]interface{})
	bodyParams := make(map[string]interface{})
	bodyParams["open_id"] = memberInfo["OpenId"]
	body["data"] = bodyParams
	params["body"] = body

	HttpRequestLib.Params = params
	res, data := HttpRequestLib.ServicePost()
	if res {
		resMap := make(map[string]interface{})
		json.Unmarshal([]byte(data), &resMap)
		authorization = resMap["authorization"].(string)
	}

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message":      "success",
		"access_token": authorization,
		"data":         memberInfo,
	})
}

// swagger:operation POST /member/register member register
// ---
// summary: 用户注册
// description: 用户注册
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: open_id
//   in: body
//   description: 用户openid
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) Register() {
}

// swagger:operation POST /member/info member info
// ---
// summary: 获取用户信息
// description: 获取用户信息
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: open_id
//   in: body
//   description: 用户openid
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) GetInfo() {
	openId := member.Base.OpenId
	if _, ok := member.Base.Req.Params["open_id"]; ok {
		if len(member.Base.Req.Params["open_id"]) > 0 {
			openId = member.Base.Req.Params["open_id"][0]
		}
	}

	MemberService := new(services.Member)
	memberInfo := MemberService.GetInfo(MemberService.GetIdByOpenId(openId))

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": memberInfo,
	})
}

func (member *Member) Ping() {
	member.Base.Res.String(http.StatusOK, "Pong")
}
