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

	jwt := new(core.Jwt)

	authorization := ""
	if len(member.Base.Req.Header["Authorization"]) > 0 {
		authorization = member.Base.Req.Header["Authorization"][0]
	}
	if !jwt.Check(authorization) {
		member.Base.Res.JSON(http.StatusUnauthorized, core.H{
			"message": "Unauthorized",
		})
	}
	fmt.Println(jwt.Get(authorization, "username").(string))
	return

	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	jwtData := jwt.New(data)
	fmt.Println(jwtData)

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
