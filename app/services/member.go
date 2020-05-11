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
	member.GetInfo(userId)
	return
}

func (member *Member) GetInfo(userId int64) {
	member.Base.Redis.NewEngine()
	key := fmt.Sprintf("member:info:%x", userId)
	fmt.Println(key)
	member.Base.Redis.Engine.SetNX(key, "Hello World!!!", time.Duration(0)*time.Second)
	fmt.Println(member.Base.Redis.Engine.Get(key))
}
