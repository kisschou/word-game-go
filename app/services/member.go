package services

import (
	"fmt"
	"strconv"
	"time"

	"wordgame/app/models"
	"wordgame/tdog/core"
)

type (
	MemberInfo map[string]interface{}

	Member struct {
		Base   core.Service
		IpAddr string
	}
)

func cacheToRedis() {
}

func (member *Member) Login(username string, password string) (memberInfo MemberInfo, err error) {
	UserModel := new(models.UserModel)
	UserModel.IpAddr = member.IpAddr
	memberInfo, err = UserModel.Login(username, password)
	if err != nil {
		return
	}

	// cache to redis.
	member.Base.Redis.NewEngine()
	userId := memberInfo["Id"]
	key := fmt.Sprintf("user:info:%x", userId)
	delete(memberInfo, "Id")
	delete(memberInfo, "Password")
	if member.Base.Redis.Engine.Exists(key).Val() > 0 {
		member.Base.Redis.Engine.Del(key)
	}

	for k, v := range memberInfo {
		member.Base.Redis.Engine.HSet(key, k, v)
	}
	key = "user:openid:" + memberInfo["OpenId"].(string)
	member.Base.Redis.Engine.SetNX(key, userId, time.Duration(0)*time.Second)
	return
}

func (member *Member) GetInfo(userId int64) (memberInfo MemberInfo) {
	member.Base.Redis.NewEngine()
	key := fmt.Sprintf("user:info:%x", userId)
	if member.Base.Redis.Engine.Exists(key).Val() > 0 {
		memberInfoMap := make(map[string]interface{})
		for k, v := range member.Base.Redis.Engine.HGetAll(key).Val() {
			memberInfoMap[k] = v
		}
		memberInfo = memberInfoMap
		return
	}

	UserModel := new(models.UserModel)
	memberInfo, err := UserModel.GetInfo(userId)
	if err != nil {
		return
	}

	// cache to redis.
	delete(memberInfo, "Id")
	delete(memberInfo, "Password")
	for k, v := range memberInfo {
		member.Base.Redis.Engine.HSet(key, k, v)
	}
	key = "user:openid:" + memberInfo["OpenId"].(string)
	member.Base.Redis.Engine.SetNX(key, userId, time.Duration(0)*time.Second)
	return
}

func (member *Member) GetIdByOpenId(openId string) (userId int64) {
	member.Base.Redis.NewEngine()
	key := "user:openid:" + openId
	if member.Base.Redis.Engine.Exists(key).Val() > 0 {
		userId, _ = strconv.ParseInt(member.Base.Redis.Engine.Get(key).Val(), 10, 64)
		return
	}

	// cannot found in redis, get from db.
	UserModel := new(models.UserModel)
	userId = UserModel.GetIdByOpenId(openId)
	return
}
