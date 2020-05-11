package models

import (
	"time"
)

type VerificationCode struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Code       string    `xorm:"comment('验证码') VARCHAR(255)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
	Status     int       `xorm:"comment('状态：1有效、0过期') BIT(1)"`
	Type       string    `xorm:"comment('验证码类型：email或者短信') VARCHAR(255)"`
	Value      string    `xorm:"comment('接收邮箱或者手机号码') VARCHAR(255)"`
	Scenes     string    `xorm:"comment('业务名称：如重置邮箱、重置密码等') VARCHAR(255)"`
}
