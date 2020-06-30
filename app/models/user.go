package models

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type (
	UserModel struct {
		Base   core.Model
		IpAddr string
	}

	User struct {
		Id                    int64     `xorm:"pk comment('ID') BIGINT(20)"`
		AvatarId              int64     `xorm:"comment('头像') index BIGINT(20)"`
		Email                 string    `xorm:"comment('邮箱') unique VARCHAR(255)"`
		Enabled               int       `xorm:"comment('状态：1启用、0禁用') TINYINT(1)"`
		Password              string    `xorm:"comment('密码') VARCHAR(255)"`
		Salt                  string    `xorm:"comment('密码加盐') VARCHAR(16)"`
		Username              string    `xorm:"comment('用户名') unique VARCHAR(255)"`
		DeptId                int64     `xorm:"comment('部门名称') index BIGINT(20)"`
		Phone                 string    `xorm:"comment('手机号码') VARCHAR(255)"`
		JobId                 int64     `xorm:"comment('岗位名称') index BIGINT(20)"`
		CreateTime            time.Time `xorm:"comment('创建日期') DATETIME"`
		LastPasswordResetTime time.Time `xorm:"comment('最后修改密码的日期') DATETIME"`
		NickName              string    `xorm:"VARCHAR(255)"`
		Sex                   string    `xorm:"VARCHAR(255)"`
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

func (userModel *UserModel) Login(username string, password string) (memberInfo map[string]interface{}, err error) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, err := false, nil

	// 检测用户名的类型
	UtilLib := new(lib.Util)
	strType := UtilLib.CheckStrType(username)

	// 用户名登入
	if strType == 0 {
		result, err = userModel.Base.Sql.Engine.Where("username=?", username).Get(userInfo)
	}

	// 邮箱登入
	if strType == 1 {
		result, err = userModel.Base.Sql.Engine.Where("email=?", username).Get(userInfo)
	}

	// 手机号登入
	if strType == 2 {
		result, err = userModel.Base.Sql.Engine.Where("phone=?", username).Get(userInfo)
	}

	if !result || userInfo.Enabled != 1 {
		userInfo = nil
		err = errors.New("ERROR_USER_LOGIN_LOCKED")
		return
	}

	CryptLib := new(lib.Crypt)
	password = CryptLib.BiuPwdBuilder(userInfo.Salt, password)
	if password != userInfo.Password {
		userInfo = nil
		err = errors.New("ERROR_USER_LOGIN_PASSWORD")
		return
	}

	// update login info.
	/*
		updateData := make(map[string]interface{})
		updateData["login"] = userInfo.Login + 1
		updateData["last_login_time"] = time.Now().Unix()
		updateData["last_login_ip"] = userModel.IpAddr
		userModel.UpdateInfo(userInfo.Id, updateData)
	*/

	memberInfo = userInfo.ToMap()
	return
}

func (userModel *UserModel) UpdateInfo(userId int64, updateData map[string]interface{}) bool {
	_, err := userModel.Base.Sql.Engine.Table(new(User)).ID(userId).Update(updateData)
	if err != nil {
		LoggerLib := new(lib.Logger)
		LoggerLib.Level = 0
		LoggerLib.Key = "error"
		LoggerLib.New(err.Error())
		return false
	}
	return true
}

func (userModel *UserModel) GetInfo(userId int64) (memberInfo map[string]interface{}, err error) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, err := userModel.Base.Sql.Engine.Where("id=?", userId).Get(userInfo)
	if !result {
		err = errors.New("ERROR_USER_LOGIN_MISSING")
		return
	}
	memberInfo = userInfo.ToMap()
	return
}

func (userModel *UserModel) Register(username, password string) (err error) {
	userInfo := new(User)
	addInfo := new(User)
	userModel.Base.Sql.NewEngine()
	var result int64

	// 检测用户名的类型
	UtilLib := new(lib.Util)
	strType := UtilLib.CheckStrType(username)

	// 用户名登入
	if strType == 0 {
		result, err = userModel.Base.Sql.Engine.Where("username=?", username).Count(userInfo)
		addInfo.Username = username
		err = errors.New("ERROR_USER_USERNAME_EXISTS")
	}

	// 邮箱登入
	if strType == 1 {
		result, err = userModel.Base.Sql.Engine.Where("email=?", username).Count(userInfo)
		addInfo.Email = username
		err = errors.New("ERROR_USER_EMAIL_EXISTS")
	}

	// 手机号登入
	if strType == 2 {
		result, err = userModel.Base.Sql.Engine.Where("phone=?", username).Count(userInfo)
		addInfo.Phone = username
		err = errors.New("ERROR_USER_PHONE_EXISTS")
	}

	if result > 0 {
		return
	}

	// 获取用户id - 雪花算法
	SnowFlakeLib := new(lib.SnowFlake)
	SnowFlakeLib.MachineId = UtilLib.GetMachineId()
	SnowFlakeLib.SN = UtilLib.RandInt64(1000, 9999)
	SnowFlakeLib.LastTimeStamp = time.Now().UnixNano() / 1000000
	userId := SnowFlakeLib.New()
	addInfo.Id = userId

	// 默认头像
	addInfo.AvatarId = 6684432467793088512

	// 获取salt和新password
	CryptLib := new(lib.Crypt)
	salt, newPassword := CryptLib.BiuPwdNewBuilder(password)
	addInfo.Salt = salt
	addInfo.Password = newPassword

	// 其他数据
	addInfo.Enabled = 1
	addInfo.NickName = "GHOST_" + strconv.FormatInt(time.Now().Unix(), 10)
	addInfo.CreateTime = time.Now()

	// 入库
	affected, err := userModel.Base.Sql.Engine.Insert(addInfo)
	if affected < 1 && err != nil {
		return
	}

	err = nil
	return
}
