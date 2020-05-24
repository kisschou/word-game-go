package models

import (
	"time"
)

type AuthDept struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Name       string    `xorm:"not null comment('名称') VARCHAR(255)"`
	Pid        int64     `xorm:"not null comment('上级部门') BIGINT(20)"`
	Enabled    int       `xorm:"not null comment('状态') BIT(1)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
}
