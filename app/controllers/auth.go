package controllers

import (
	"net/http"

	"wordgame/tdog/core"
)

type Auth struct {
	Base core.Controller
}

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
