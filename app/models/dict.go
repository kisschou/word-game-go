package models

import (
	"time"
)

type Dict struct {
	Id         int64     `xorm:"pk autoincr BIGINT(11)"`
	Name       string    `xorm:"not null comment('字典名称') VARCHAR(255)"`
	Remark     string    `xorm:"comment('描述') VARCHAR(255)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
}
