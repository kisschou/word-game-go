package models

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
	. "wordgame/library/cache"
	. "wordgame/library/database"
	. "wordgame/library/encrypt"
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
func updateLoginInfo(memberInfo *User) {
	key := "session:user:"
	key += memberInfo.OpenId
	sid := SessionId()
	token := buildToken(sid)

	// 清空同一个会话下的其他用户的登录信息
}

// 用户登录
// 1. 通过用户名或手机号或邮箱获取用户信息, 并判断用户是否存在或锁定, 及密码是否正确
// 2. 用户信息完整性检测及填充
// 3. 更新登录信息
func (info *User) Login(username string, password string) (memberInfo *User, err error) {
	memberInfo = new(User)

	result, err := Engine.Where("username=?", username).Get(memberInfo)
	if !result {
		result, err = Engine.Where("mobile=?", username).Get(memberInfo)
	}
	if !result {
		result, err = Engine.Where("emial=?", username).Get(memberInfo)
	}
	if !result || memberInfo.Status == 0 {
		memberInfo = nil
		return
	}

	if ThinkUcenterMd5(password, "sot31OWgWOFpUXA0gKQ6aIi3Y5iQ9LiRS8sKGWlVdJx9ca93SgOuSGGf3ygEYqPB", memberInfo.Salt) != memberInfo.Password {
		memberInfo = nil
		return
	}

	if len(memberInfo.Nickname) < 1 {
		_, _ = Engine.Where("id=?", memberInfo.Id).Update(&User{Nickname: memberInfo.Username})
	}

	updateLoginInfo(memberInfo)
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
