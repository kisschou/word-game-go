package services

import (
	"fmt"
	"time"

	"wordgame/app/models"
	"wordgame/tdog/core"
)

type (
	MemberInfo map[string]interface{}

	Member struct {
		Base core.Service
	}
)

func (member *Member) Login(username string, password string) (memberInfo MemberInfo, err error) {
	UserModel := new(models.UserModel)
	userId, err := UserModel.Login(username, password)
	if err != nil {
		return
	}
	return
}

func (member *Member) GetInfo(userId string) {
	member.Base.Redis.NewEngine()
	member.Base.Redis.Engine.SetNX("member:uid:1", "Hello World!!!", time.Duration(0)*time.Second)
	fmt.Println(member.Base.Redis.Engine.Get("member:info:" + userId))
}
