package models

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	_ "strings"
	"time"
	. "wordgame/library/cache"
	. "wordgame/library/database"
	. "wordgame/library/encrypt"
	. "wordgame/library/unit"
)

type User struct {
	Id                int32   `json:"id" notnull pk binding:"required"`
	OpenId            string  `json:"open_id" binding:"required"`
	Username          string  `json:"username" binding:"required"`
	Password          string  `json:"password" binding:"required"`
	DrawalPassword    string  `json:"drawal_password" binding:"required"`
	Salt              string  `json:"salt" binding:"required"`
	Email             string  `json:"email" binding:"required"`
	IsEmailVerify     int32   `json:"is_email_verify" binding:"required"`
	Mobile            string  `json:"mobile" binding:"required"`
	IsMobileVerify    int32   `json:"is_mobile_verify" binding:"required"`
	AvatarId          int32   `json:"avatar_id" binding:"required"`
	CoverId           int32   `json:"cover_id" binding:"required"`
	InviteUserId      int32   `json:"invite_user_id" binding:"required"`
	Realname          string  `json:"realname" binding:"required"`
	Nickname          string  `json:"nickname" binding:"required"`
	Firstname         string  `json:"firstname" binding:"required"`
	Lastname          string  `json:"lastname" binding:"required"`
	Birthday          int32   `json:"birthday" binding:"required"`
	Sex               int32   `json:"sex" binding:"required"`
	CurrencyId        int32   `json:"currency_id" binding:"required"`
	Money             float32 `json:"money" binding:"required"`
	FreezeMoney       float32 `json:"freeze_money" binding:"required"`
	Score             int32   `json:"score" binding:"required"`
	PowerScore        int32   `json:"power_score" binding:"required"`
	ScalperOrderCount int32   `json:"scalper_order_count" binding:"required"`
	Address           string  `json:"address" binding:"required"`
	Longitude         int32   `json:"longitude" binding:"required"`
	Latitude          int32   `json:"latitude" binding:"required"`
	ProvinceId        int32   `json:"province_id" binding:"required"`
	CityId            int32   `json:"city_id" binding:"required"`
	IsIdentity        int32   `json:"is_identity" binding:"required"`
	IdCardNumber      string  `json:"id_card_number" binding:"required"`
	IdCardFront       int32   `json:"id_card_front" binding:"required"`
	IdCardBack        int32   `json:"id_card_back" binding:"required"`
	RegIp             int64   `json:"reg_ip" binding:"required"`
	RegTime           int32   `json:"reg_time" binding:"required"`
	LastLoginIp       int64   `json:"last_login_ip" binding:"required"`
	LastLoginTime     int32   `json:"last_login_time" binding:"required"`
	ForbiddenReason   string  `json:"forbidden_reason" binding:"required"`
	UpdateTime        int32   `json:"update_time" binding:"required"`
	Status            int32   `json:"status" binding:"required"`
}

func buildToken(sid string) string {
	sid += "sot31OWgWOFpUXA0gKQ6aIi3Y5iQ9LiRS8sKGWlVdJx9ca93SgOuSGGf3ygEYqPB"
	sid += strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	sid += strconv.Itoa(rand.Intn(9999))
	return Md5(sid)
}

// 更新用户信息
func updateLoginInfo(memberInfo *User) map[string]interface{} {
	sid := SessionId()
	token := buildToken(sid)

	// 删除前者登录数据
	key := "session:user:"
	key += memberInfo.OpenId
	if Redis.Exists(key).Val() == 1 {
		oldToken := Redis.HGet(key, "token").Val()
		key = "session:"
		key += oldToken
		Redis.Del(key)
	}

	// 记录Token和OpenId的关联到Redis
	key = "session:"
	key += token
	// Redis.SetNX(key, memberInfo.OpenId, time.Duration(120)*time.Second)
	Redis.Set(key, memberInfo.OpenId, 0)

	// 记录用户信息到Redis
	key = "session:user:"
	key += memberInfo.OpenId

	// Struct2Map
	memberInfoMap := Struct2Map(memberInfo)
	memberInfoMap["session_id"] = sid
	memberInfoMap["token"] = token

	for k, v := range memberInfoMap {
		Redis.HMSet(key, k, v)
	}
	/*
		if strings.Contains(Redis.TTL(key).Val().String(), "ns") {
			Redis.Expire(key, time.Duration(86400)*time.Second)
		}
	*/

	return memberInfoMap
}

// session判断规则:
// 说明:
// session:{:token} - 表示当前未过期的的session value:string({:openId}) ttl:-1
// session:user:{:openId} - 表示已经缓存的用户信息列表 value:hash({用户信息k-v}) ttl:86400s
// 判断流程:
// 1. 判断session:{:token}是否存在
// 2. 取出{:openId}判断session:user:{:openId}是否存在
// 返回:
// 判断(1)错误发生时 如果请求没有携带openId则判定为 未登录
// 判断(2)错误发生时 如果session存在, 则自动缓存用户信息
func (info *User) UAuth(token string) (err error) {
	if len(token) < 1 {
		err = errors.New("用户未登录")
		return
	}

	key := "session:"
	key += token
	// 未登录
	if Redis.Exists(key).Val() != 1 {
		err = errors.New("用户未登录")
		return
	}

	openId := Redis.Get(key).Val()
	key = "session:user:"
	key += openId
	if Redis.Exists(key).Val() != 1 {
		// 缓存用户信息
		info.OpenId = openId
		_, err = info.GetInfo()
		if err != nil {
			return
		}
	} else {
		userToken := Redis.HGet(key, "token").Val()
		if userToken != token {
			key = "session:"
			key += token
			Redis.Del(key)
			err = errors.New("用户未登录")
			return
		}
	}

	return
}

func (info *User) GetInfo() (memberInfoMap map[string]interface{}, err error) {
	openId := info.OpenId
	if len(openId) < 1 {
		err = errors.New("未找到相关用户数据")
		return
	}

	// 1. 判断redis中是否存在
	key := "session:user:"
	key += openId
	if Redis.Exists(key).Val() == 1 {
		data := Redis.HGetAll(key).Val()
		j, _ := json.Marshal(data)
		json.Unmarshal(j, &memberInfoMap)
		delete(memberInfoMap, "session_id")
		delete(memberInfoMap, "token")
		return
	}

	// 2. 从数据库中获取
	memberInfo := new(User)
	result, err := Engine.Where("open_id=?", openId).Get(memberInfo)
	if !result {
		err = errors.New("未找到相关用户数据")
		return
	}

	// 3. 缓存到redis
	memberInfoMap = Struct2Map(memberInfo)
	for k, v := range memberInfoMap {
		Redis.HMSet(key, k, v)
	}
	/*
		if strings.Contains(Redis.TTL(key).Val().String(), "ns") {
			Redis.Expire(key, time.Duration(86400)*time.Second)
		}
	*/

	return
}

// 用户登录
// 1. 通过用户名或手机号或邮箱获取用户信息, 并判断用户是否存在或锁定, 及密码是否正确
// 2. 用户信息完整性检测及填充
// 3. 更新登录信息
func (info *User) Login() (memberInfoMap map[string]interface{}, err error) {
	username := info.Username
	password := info.Password
	if len(username) < 1 || len(password) < 6 {
		err = errors.New("参数错误")
		return
	}

	memberInfo := new(User)

	result, err := Engine.Where("username=?", username).Get(memberInfo)
	if !result {
		result, err = Engine.Where("mobile=?", username).Get(memberInfo)
	}
	if !result {
		result, err = Engine.Where("emial=?", username).Get(memberInfo)
	}
	if !result || memberInfo.Status == 0 {
		err = errors.New("用户不存在或者已经锁定!")
		return
	}

	if ThinkUcenterMd5(password, "sot31OWgWOFpUXA0gKQ6aIi3Y5iQ9LiRS8sKGWlVdJx9ca93SgOuSGGf3ygEYqPB", memberInfo.Salt) != memberInfo.Password {
		err = errors.New("密码错误!")
		return
	}

	if len(memberInfo.Nickname) < 1 {
		_, _ = Engine.Where("id=?", memberInfo.Id).Update(&User{Nickname: memberInfo.Username})
	}

	memberInfoMap = updateLoginInfo(memberInfo)
	return
}

func (info *User) Insert() (userId uint, err error) {
	result, err := Engine.Insert(&info)
	if result != 0 {
		return
	}
	return
}

func (info *User) FindAll() (memberList []User, err error) {
	result := Engine.Find(&memberList)

	if result != nil {
		err = result
		return
	}

	return
}
