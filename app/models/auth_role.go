package models

import (
	"time"
)

type AuthRole struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Name       string    `xorm:"not null comment('名称') VARCHAR(255)"`
	Remark     string    `xorm:"comment('备注') VARCHAR(255)"`
	DataScope  string    `xorm:"comment('数据权限') VARCHAR(255)"`
	Level      int       `xorm:"comment('角色级别') INT(255)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
	Permission string    `xorm:"comment('功能权限') VARCHAR(255)"`
}
