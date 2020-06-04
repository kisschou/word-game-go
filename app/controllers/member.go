package controllers

import (
	"net/http"

	"wordgame/app/services"
	"wordgame/tdog/core"
)

type Member struct {
	Base core.Controller
}

// swagger:operation POST /member/login
// ---
// summary: 用户登录
// description: 用户登录
// parameters:
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
	memberInfo, err := MemberService.Login(username, password)
	if err != nil {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	jwt := new(core.Jwt)
	member.Base.Res.JSON(http.StatusOK, core.H{
		"message":      "success",
		"access_token": jwt.New(map[string]interface{}{"open_id": memberInfo["OpenId"]}),
		"data":         memberInfo,
	})
}

// swagger:operation POST /member/info
// ---
// summary: 获取用户信息
// description: 获取用户信息
// parameters:
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
