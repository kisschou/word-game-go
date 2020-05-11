package models

import (
	"time"
)

type AuthJob struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Name       string    `xorm:"not null comment('岗位名称') VARCHAR(255)"`
	Enabled    int       `xorm:"not null comment('岗位状态') BIT(1)"`
	Sort       int64     `xorm:"not null comment('岗位排序') BIGINT(20)"`
	DeptId     int64     `xorm:"comment('部门ID') index BIGINT(20)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
}
