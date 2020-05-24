package models

import (
	"errors"
	"reflect"
	"time"

	"wordgame/tdog/core"
)

type (
	UserModel struct {
		Base core.Model
	}

	User struct {
		Id                    int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
		OpenId                string    `xorm:"comment('用户openid') unique CHAR(40)"`
		AvatarId              int64     `xorm:"comment('头像') index BIGINT(20)"`
		Email                 string    `xorm:"comment('邮箱') unique VARCHAR(255)"`
		Enabled               int64     `xorm:"comment('状态：1启用、0禁用') BIGINT(20)"`
		Password              string    `xorm:"comment('密码') VARCHAR(255)"`
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

func (userModel *UserModel) Login(username string, password string) (userId int64, err error) {
	userInfo := new(User)
	userModel.Base.Sql.NewEngine()
	result, err := userModel.Base.Sql.Engine.Where("username=?", username).Get(userInfo)
	if !result || userInfo.Enabled != 1 {
		userInfo = nil
		err = errors.New("用户不存在或者已经锁定!")
		return
	}
	if password != userInfo.Password {
		userInfo = nil
		err = errors.New("密码错误!")
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
		err = errors.New("未找到该用户数据!")
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
