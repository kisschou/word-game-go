package controllers

import (
	"net/http"

	"github.com/kisschou/tdog"
)

type Auth struct {
	Base tdog.Controller
}

func (auth *Auth) GetToken() {
	jwt := new(tdog.Jwt)
	authorization := jwt.New(make(map[string]interface{}))
	if _, ok := auth.Base.Req.Header["Authorization"]; ok {
		if len(auth.Base.Req.Header["Authorization"]) > 0 && len(auth.Base.Req.Header["Authorization"][0]) > 0 {
			authorization = jwt.Refresh(auth.Base.Req.Header["Authorization"][0])
		}
	}
	auth.Base.Res.JSON(http.StatusOK, tdog.H{
		"authorization": authorization,
	})
}
