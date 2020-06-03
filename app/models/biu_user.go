package models

import (
	"errors"
	"reflect"

	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type (
	UserModel struct {
		Base core.Model
	}

	User struct {
		Id                int64  `xorm:"pk comment('用户ID') BIGINT(20)"`
		OpenId            string `xorm:"not null comment('用户唯一标识,供返回客户端使用') CHAR(16)"`
		Username          string `xorm:"not null comment('用户名') unique CHAR(32)"`
		Password          string `xorm:"not null default '' comment('密码') CHAR(32)"`
		DrawalPassword    string `xorm:"not null default '' comment('提现密码') CHAR(32)"`
		Salt              string `xorm:"not null default '' comment('密码随机码') CHAR(6)"`
		Email             string `xorm:"not null comment('邮箱') unique CHAR(32)"`
		IsEmailVerify     int    `xorm:"not null default 0 comment('邮箱是否已经验证') TINYINT(1)"`
		Mobile            string `xorm:"not null comment('用户手机') unique CHAR(16)"`
		IsMobileVerify    int    `xorm:"not null default 0 comment('手机是否已经验证') TINYINT(1)"`
		AvatarId          int    `xorm:"not null default 0 comment('头像ID') INT(10)"`
		CoverId           int    `xorm:"not null default 0 comment('封面ID') INT(10)"`
		InviteUserId      int    `xorm:"not null default 0 comment('邀请注册的用户ID') INT(10)"`
		Realname          string `xorm:"not null default '' comment('真实姓名') CHAR(32)"`
		Nickname          string `xorm:"not null default '' comment('昵称') CHAR(32)"`
		Firstname         string `xorm:"not null default '' comment('英文名') CHAR(32)"`
		Lastname          string `xorm:"not null default '' comment('英文姓') CHAR(32)"`
		Birthday          int    `xorm:"not null default 0 comment('生日') INT(10)"`
		Sex               int    `xorm:"not null default 0 comment('性别') TINYINT(3)"`
		CurrencyId        int    `xorm:"not null default 0 comment('账户默认币种ID') INT(10)"`
		Money             string `xorm:"not null default 0.0000 comment('余额') DECIMAL(20,4)"`
		FreezeMoney       string `xorm:"not null default 0.0000 comment('冻结金额') DECIMAL(20,4)"`
		Score             int    `xorm:"not null default 0 comment('用户积分') INT(10)"`
		PowerScore        int    `xorm:"not null default 0 comment('app积分') INT(10)"`
		ScalperOrderCount int    `xorm:"not null default 0 comment('app总单数') INT(10)"`
		Address           string `xorm:"not null default '' comment('地址') CHAR(64)"`
		Longitude         int    `xorm:"not null default 0 comment('经度') INT(10)"`
		Latitude          int    `xorm:"not null default 0 comment('纬度') INT(10)"`
		ProvinceId        int    `xorm:"not null default 0 comment('省份id') INT(10)"`
		CityId            int    `xorm:"not null default 0 comment('城市id') INT(10)"`
		IsIdentity        int    `xorm:"not null default 0 comment('是否实名认证') TINYINT(1)"`
		IdCardNumber      string `xorm:"not null default '' comment('身份证号码') VARCHAR(32)"`
		IdCardFront       int    `xorm:"not null default 0 comment('身份证正面图片ID') INT(10)"`
		IdCardBack        int    `xorm:"not null default 0 comment('身份证背面图片ID') INT(10)"`
		Login             int    `xorm:"not null default 0 comment('登录次数') INT(10)"`
		RegIp             int64  `xorm:"not null default 0 comment('注册IP') BIGINT(20)"`
		RegTime           int64  `xorm:"not null default 0 comment('注册时间') BIGINT(20)"`
		LastLoginIp       int64  `xorm:"not null default 0 comment('最后登录IP') BIGINT(20)"`
		LastLoginTime     int64  `xorm:"not null default 0 comment('最后登录时间') BIGINT(20)"`
		ForbiddenReason   string `xorm:"not null default '' comment('禁用原因') VARCHAR(255)"`
		UpdateTime        int64  `xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
		Status            int    `xorm:"not null default 0 comment('会员状态') index TINYINT(4)"`
	}
)

func (user *User) ToMap() map[string]interface{} {
	mapVal := make(map[string]interface{})
	elem := reflect.ValueOf(user).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		mapVal[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return mapVal
}

func (userModel *UserModel) Login(username string, password string) (userId int64, err error) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, err := userModel.Base.Sql.Engine.Where("username=?", username).Get(userInfo)
	if !result || userInfo.Status != 1 {
		userInfo = nil
		err = errors.New("ERROR_LOGIN_LOCKED")
		return
	}

	CryptLib := new(lib.Crypt)
	password = CryptLib.BiuPwdBuilder(userInfo.Salt, password)
	if password != userInfo.Password {
		userInfo = nil
		err = errors.New("ERROR_LOGIN_PASSWORD")
		return
	}
	userId = userInfo.Id
	return
}

func (userModel *UserModel) GetInfo(userId int64) (memberInfo map[string]interface{}, err error) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, err := userModel.Base.Sql.Engine.Where("id=?", userId).Get(userInfo)
	if !result {
		err = errors.New("ERROR_LOGIN_MISSING")
		return
	}
	memberInfo = userInfo.ToMap()
	return
}

func (userModel *UserModel) GetIdByOpenId(openId string) (userId int64) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, _ := userModel.Base.Sql.Engine.Where("open_id=?", openId).Get(userInfo)
	if !result {
		userId = int64(0)
		return
	}
	userId = userInfo.Id
	return
}
