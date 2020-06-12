package controllers

import (
	"encoding/json"
	"net/http"

	"wordgame/tdog/core"
)

type Auth struct {
	Base core.Controller
}

// swagger:operation GET /auth/get auth get
// ---
// summary: 获取授权信息
// description: 获取授权信息
// parameters:
// - name: data
//   in: body
//   description: 数据集合
//   type: json
//   required: false
// responses:
//   200: repoResp
//   401: errMsg
func (auth *Auth) GetToken() {
	JwtCore := new(core.Jwt)
	authorization := JwtCore.New(make(map[string]interface{}))
	if _, ok := auth.Base.Req.Params["data"]; ok {
		if len(auth.Base.Req.Params["data"]) > 0 && len(auth.Base.Req.Params["data"][0]) > 0 {
			dataMap := make(map[string]interface{})
			json.Unmarshal([]byte(auth.Base.Req.Params["data"][0]), &dataMap)
			authorization = JwtCore.New(dataMap)
		}
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"authorization": authorization,
	})
}

// swagger:operation GET /auth/refresh auth refresh
// ---
// summary: 刷新授权信息
// description: 刷新授权信息
// parameters:
// - name: authorization
//   in: body
//   description: 授权信息
//   type: string
//   required: true
// responses:
//   200: repoResp
//   500: errMsg
func (auth *Auth) RefreshToken() {
	JwtCore := new(core.Jwt)
	authorization := ""
	if _, ok := auth.Base.Req.Params["authorization"]; ok {
		if len(auth.Base.Req.Params["authorization"]) > 0 && len(auth.Base.Req.Params["authorization"][0]) > 0 {
			authorization = JwtCore.Refresh(auth.Base.Req.Params["authorization"][0])
		}
	}
	if len(authorization) < 1 {
		auth.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_REQUEST_PARAMS",
		})
		return
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"authorization": authorization,
	})
}

// swagger:operation POST /auth/verify auth verify
// ---
// summary: 校验授权信息
// description: 校验授权信息
// parameters:
// - name: authorization
//   in: body
//   description: 授权信息
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (auth *Auth) VerifyToken() {
	JwtCore := new(core.Jwt)
	authorization := ""
	if _, ok := auth.Base.Req.Params["authorization"]; ok {
		if len(auth.Base.Req.Params["authorization"]) > 0 && len(auth.Base.Req.Params["authorization"][0]) > 0 {
			authorization = auth.Base.Req.Params["authorization"][0]
		}
	}
	if !JwtCore.Check(authorization) {
		auth.Base.Res.JSON(http.StatusUnauthorized, core.H{
			"code": "ERROR_UNAUTHOZED",
		})
		return
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"message": "成功",
	})
}

// swagger:operation POST /auth/verifyLogin auth verifyLogin
// ---
// summary: 校验登入状态
// description: 校验登入状态
// parameters:
// - name: authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
//   510: errMsg
func (auth *Auth) VerifyLogin() {
	JwtCore := new(core.Jwt)
	authorization := auth.Base.Req.Header["Authorization"][0]
	openId := ""
	if JwtCore.Get(authorization, "open_id") != nil {
		openId = JwtCore.Get(authorization, "open_id").(string)
	}
	if len(openId) < 1 {
		auth.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": "ERROR_UNLOGIN",
		})
		return
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"message": "成功",
		"open_id": openId,
	})
}

// swagger:operation POST /auth/getKey auth getKey
// ---
// summary: 获取指定数据
// description: 获取指定数据
// parameters:
// - name: authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// - name: key
//   in: body
//   description: 获取键名
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (auth *Auth) GetKey() {
	JwtCore := new(core.Jwt)
	authorization := auth.Base.Req.Header["Authorization"][0]
	key := ""
	value := ""
	if _, ok := auth.Base.Req.Params["key"]; ok {
		if len(auth.Base.Req.Params["key"]) > 0 && len(auth.Base.Req.Params["key"][0]) > 0 {
			key = auth.Base.Req.Params["key"][0]
		}
	}
	if JwtCore.Get(authorization, key) != nil {
		value = JwtCore.Get(authorization, key).(string)
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"message": "获取成功",
		"key":     key,
		"value":   value,
	})
}

// swagger:operation POST /auth/getAllKeys auth getAllKeys
// ---
// summary: 获取所有数据
// description: 获取所有数据
// parameters:
// - name: authorization
//   in: header
//   description: 授权信息
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: errMsg
func (auth *Auth) GetAllKeys() {
	JwtCore := new(core.Jwt)
	authorization := auth.Base.Req.Header["Authorization"][0]
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"message": "获取成功",
		"data":    JwtCore.GetData(authorization),
	})
}
