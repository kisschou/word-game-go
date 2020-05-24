package models

import (
	"time"
)

type MntDatabase struct {
	Id         string    `xorm:"not null pk comment('编号') VARCHAR(255)"`
	Name       string    `xorm:"not null comment('名称') VARCHAR(255)"`
	JdbcUrl    string    `xorm:"not null comment('jdbc连接') VARCHAR(255)"`
	UserName   string    `xorm:"not null comment('账号') VARCHAR(255)"`
	Pwd        string    `xorm:"not null comment('密码') VARCHAR(255)"`
	CreateTime time.Time `xorm:"DATETIME"`
}
