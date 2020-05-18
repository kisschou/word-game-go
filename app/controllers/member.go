package controllers

import (
	"fmt"
	"net/http"

	"wordgame/app/services"
	"wordgame/tdog/core"
)

type Member struct {
	Base core.Controller
}

func (member *Member) Login() {
	username := member.Base.Req.Params["username"][0]
	password := member.Base.Req.Params["password"][0]
	fmt.Println("Login with username: [" + username + "] password: [" + password + "]")

	MemberService := new(services.Member)

	memberInfo, err := MemberService.Login(username, password)
	if err != nil {
		member.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println("LoginMemberInfo: ", memberInfo)
	member.Base.Res.JSON(http.StatusOK, core.H{
		"message": "login success",
	})
}

func (member *Member) GetInfo() {
}

func (member *Member) Ping() {
	member.Base.Res.String(http.StatusOK, "Pong")
}
