package models

import (
	"time"
)

type UserAvatar struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	RealName   string    `xorm:"comment('真实文件名') VARCHAR(255)"`
	Path       string    `xorm:"comment('路径') VARCHAR(255)"`
	Size       string    `xorm:"comment('大小') VARCHAR(255)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
}
