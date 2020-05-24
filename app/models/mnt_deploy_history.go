package models

import (
	"time"
)

type MntDeployHistory struct {
	Id         string    `xorm:"not null pk comment('编号') VARCHAR(50)"`
	AppName    string    `xorm:"not null comment('应用名称') VARCHAR(255)"`
	DeployDate time.Time `xorm:"not null comment('部署日期') DATETIME"`
	DeployUser string    `xorm:"not null comment('部署用户') VARCHAR(50)"`
	Ip         string    `xorm:"not null comment('服务器IP') VARCHAR(20)"`
	DeployId   int64     `xorm:"comment('部署编号') BIGINT(20)"`
}
