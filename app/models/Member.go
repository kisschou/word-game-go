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

	"github.com/golang/protobuf/proto"
	"wordgame/data/rpc/golang/api"
	"wordgame/data/rpc/golang/base"
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

func buildProto(memberInfoMap map[string]interface{}) uint8 {
	r := BiuRpc_Base.Response{}
	r.SessionId = ""
	r.RequestId = ""
	r.Token = ""
	r.Nonce = ""
	r.Timestamp = 0
	r.Code = 0
	r.Msg = ""

	userInfo := BiuRpc_Base.User{}
	userInfo.UserId = ""
	userInfo.Username = ""
	userInfo.Nickname = ""
	userInfo.Email = ""
	userInfo.Avatar = nil
	userInfo.Mobile = ""
	userInfo.Realname = ""
	userInfo.Sex = 0
	userInfo.Birthday = 0
	userInfo.UnreadMessageCount = 0
	userInfo.IsCertificate = 0
	userInfo.IdCardNumber = ""
	userInfo.IsEmailVerify = false
	userInfo.IsMobileVerify = false
	userInfo.Firstname = ""
	userInfo.Lastname = ""
	/*
		UserId               string          `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
		Username             string          `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
		Nickname             string          `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
		Email                string          `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
		Avatar               *Picture        `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
		Mobile               string          `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
		Realname             string          `protobuf:"bytes,7,opt,name=realname,proto3" json:"realname,omitempty"`
		Sex                  SignUserSex     `protobuf:"varint,8,opt,name=sex,proto3,enum=BiuRpc.Base.SignUserSex" json:"sex,omitempty"`
		Birthday             int32           `protobuf:"varint,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
		UnreadMessageCount   int32           `protobuf:"varint,10,opt,name=unreadMessageCount,proto3" json:"unreadMessageCount,omitempty"`
		IsCertificate        int32           `protobuf:"varint,11,opt,name=isCertificate,proto3" json:"isCertificate,omitempty"`
		IdCardNumber         string          `protobuf:"bytes,12,opt,name=idCardNumber,proto3" json:"idCardNumber,omitempty"`
		IsEmailVerify        bool            `protobuf:"varint,13,opt,name=isEmailVerify,proto3" json:"isEmailVerify,omitempty"`
		IsMobileVerify       bool            `protobuf:"varint,14,opt,name=isMobileVerify,proto3" json:"isMobileVerify,omitempty"`
		Firstname            string          `protobuf:"bytes,15,opt,name=firstname,proto3" json:"firstname,omitempty"`
		Lastname             string          `protobuf:"bytes,16,opt,name=lastname,proto3" json:"lastname,omitempty"`
		CurrencyId           int32           `protobuf:"varint,17,opt,name=currencyId,proto3" json:"currencyId,omitempty"`
		RegTime              int32           `protobuf:"varint,18,opt,name=regTime,proto3" json:"regTime,omitempty"`
		LastLoginTime        int32           `protobuf:"varint,19,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
		Money                string          `protobuf:"bytes,20,opt,name=money,proto3" json:"money,omitempty"`
		FreezeMoney          string          `protobuf:"bytes,21,opt,name=freezeMoney,proto3" json:"freezeMoney,omitempty"`
		PowerScore           int32           `protobuf:"varint,22,opt,name=powerScore,proto3" json:"powerScore,omitempty"`
		ScalperOrderCount    int32           `protobuf:"varint,23,opt,name=scalperOrderCount,proto3" json:"scalperOrderCount,omitempty"`
		InterestList         []*UserInterest `protobuf:"bytes,24,rep,name=interestList,proto3" json:"interestList,omitempty"`
		AmazonAccount        string          `protobuf:"bytes,25,opt,name=amazonAccount,proto3" json:"amazonAccount,omitempty"`
		AmazonEmail          string          `protobuf:"bytes,26,opt,name=amazonEmail,proto3" json:"amazonEmail,omitempty"`
		HasDrawalAccount     bool            `protobuf:"varint,27,opt,name=hasDrawalAccount,proto3" json:"hasDrawalAccount,omitempty"`
		HasDrawalPwd         bool            `protobuf:"varint,28,opt,name=hasDrawalPwd,proto3" json:"hasDrawalPwd,omitempty"`
	*/

	res := BiuRpc_Api.ResponseUserLogin{}
	res.Base = &r
	res.SessionId = ""
	res.SessionType = 0
	res.Token = ""
	res.UserInfo = &userInfo
	/*
		Base                 *Response `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
		SessionId            string    `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
		SessionType          int32     `protobuf:"varint,3,opt,name=sessionType,proto3" json:"sessionType,omitempty"`
		Token                string    `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
		UserInfo             *User     `protobuf:"bytes,5,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	*/
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
