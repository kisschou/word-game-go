package service

import (
	"fmt"
	"time"

	"wordgame/app/model"
	"wordgame/tdog/core"
)

type (
	MemberInfo map[string]interface{}

	Member struct {
		Base core.Service
	}
)

func (member *Member) Login(username string, password string) {
	MemberModel := new(model.Member)
	MemberModel.Login(username, password)
}

func (member *Member) GetInfo(userId string) {
	member.Base.Redis.NewEngine()
	member.Base.Redis.Engine.SetNX("member:uid:1", "Hello World!!!", time.Duration(0)*time.Second)
	fmt.Println(member.Base.Redis.Engine.Get("member:info:" + userId))
}
