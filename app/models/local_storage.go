package models

import (
	"time"
)

type LocalStorage struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	RealName   string    `xorm:"comment('文件真实的名称') VARCHAR(255)"`
	Name       string    `xorm:"comment('文件名') VARCHAR(255)"`
	Suffix     string    `xorm:"comment('后缀') VARCHAR(255)"`
	Path       string    `xorm:"comment('路径') VARCHAR(255)"`
	Type       string    `xorm:"comment('类型') VARCHAR(255)"`
	Size       string    `xorm:"comment('大小') VARCHAR(100)"`
	Operate    string    `xorm:"comment('操作人') VARCHAR(255)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
}
