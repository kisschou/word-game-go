package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	bodyParams["user_id"] = memberInfo["Id"]
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
//   401: errMsg
func (member *Member) Register() {
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
	err := MemberService.Register(username, password)

	if err != nil {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
	})
}

// swagger:operation POST /member/updateInfo member updateInfo
// ---
// summary: 更新用户信息
// description: 更新用户信息
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: avatar_id
//   in: body
//   description: 头像id
//   type: string
//   required: false
// - name: email
//   in: body
//   description: 邮箱
//   type: string
//   required: false
// - name: phone
//   in: body
//   description: 手机号码
//   type: string
//   required: false
// - name: nick_name
//   in: body
//   description: 昵称
//   type: string
//   required: false
// - name: sex
//   in: body
//   description: 性别
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) UpdateInfo() {
	userId := member.Base.UserId
	updateInfo := make(map[string]interface{})
	if _, ok := member.Base.Req.Params["avatar_id"]; ok {
		if len(member.Base.Req.Params["avatar_id"]) > 0 {
			updateInfo["avatar_id"] = member.Base.Req.Params["avatar_id"][0]
		}
	}
	if _, ok := member.Base.Req.Params["email"]; ok {
		if len(member.Base.Req.Params["email"]) > 0 {
			updateInfo["email"] = member.Base.Req.Params["email"][0]
		}
	}
	if _, ok := member.Base.Req.Params["phone"]; ok {
		if len(member.Base.Req.Params["phone"]) > 0 {
			updateInfo["phone"] = member.Base.Req.Params["phone"][0]
		}
	}
	if _, ok := member.Base.Req.Params["nick_name"]; ok {
		if len(member.Base.Req.Params["nick_name"]) > 0 {
			updateInfo["nick_name"] = member.Base.Req.Params["nick_name"][0]
		}
	}
	if _, ok := member.Base.Req.Params["sex"]; ok {
		if len(member.Base.Req.Params["sex"]) > 0 {
			updateInfo["sex"] = member.Base.Req.Params["sex"][0]
		}
	}

	if len(updateInfo) < 1 {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	MemberService := new(services.Member)
	if !MemberService.UpdateInfo(userId, updateInfo) {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_USER_UPDATE_FAIL",
		})
		return
	}

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
	})
}

// swagger:operation POST /member/changePassword member changePassword
// ---
// summary: 修改密码
// description: 修改密码
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: password
//   in: body
//   description: 新密码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) ChangePassword() {
	// userId := member.Base.UserId
	password := ""
	if _, ok := member.Base.Req.Params["password"]; ok {
		if len(member.Base.Req.Params["password"]) > 0 {
			password = member.Base.Req.Params["password"][0]
		}
	}

	if len(password) < 6 {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
	})
}

// swagger:operation POST /member/resetPassword member resetPassword
// ---
// summary: 重设密码-忘记密码
// description: 重设密码-忘记密码
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: false
// - name: password
//   in: body
//   description: 新密码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) ResetPassword() {
}

// swagger:operation POST /member/initPassword member initPassword
// ---
// summary: 初始化密码
// description: 初始化密码
// parameters:
// - name: Authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: password
//   in: body
//   description: 新密码
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) InitPassword() {
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
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (member *Member) GetInfo() {
	userId := member.Base.UserId
	if _, ok := member.Base.Req.Params["user_id"]; ok {
		if len(member.Base.Req.Params["user_id"]) > 0 {
			var err error
			userId, err = strconv.ParseInt(member.Base.Req.Params["user_id"][0], 10, 64)
			if err != nil {
				userId = member.Base.UserId
			}
		}
	}

	MemberService := new(services.Member)
	memberInfo := MemberService.GetInfo(userId)

	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": memberInfo,
	})
}

func (member *Member) Ping() {
	member.Base.Res.String(http.StatusOK, "Pong")
}
