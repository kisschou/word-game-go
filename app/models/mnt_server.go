package models

import (
	"time"
)

type MntServer struct {
	Id         int64     `xorm:"pk autoincr comment('IP地址') BIGINT(20)"`
	Account    string    `xorm:"comment('账号') VARCHAR(255)"`
	Ip         string    `xorm:"comment('IP地址') VARCHAR(255)"`
	Name       string    `xorm:"comment('名称') VARCHAR(255)"`
	Password   string    `xorm:"comment('密码') VARCHAR(255)"`
	Port       int       `xorm:"comment('端口') INT(11)"`
	CreateTime time.Time `xorm:"DATETIME"`
}
