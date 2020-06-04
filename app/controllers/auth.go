package controllers

import (
	"net/http"

	"wordgame/tdog/core"
)

type Auth struct {
	Base core.Controller
}

// swagger:operation POST /auth/getToken
// ---
// summary: 获取授权信息
// description: 获取授权信息
// parameters:
// responses:
//   200: repoResp
//   401: errMsg
func (auth *Auth) GetToken() {
	jwt := new(core.Jwt)
	authorization := jwt.New(make(map[string]interface{}))
	if _, ok := auth.Base.Req.Header["Authorization"]; ok {
		if len(auth.Base.Req.Header["Authorization"]) > 0 && len(auth.Base.Req.Header["Authorization"][0]) > 0 {
			authorization = jwt.Refresh(auth.Base.Req.Header["Authorization"][0])
		}
	}
	auth.Base.Res.JSON(http.StatusOK, core.H{
		"authorization": authorization,
	})
}
