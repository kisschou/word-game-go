package models

import (
	"time"
)

type DictDetail struct {
	Id         int64     `xorm:"pk autoincr BIGINT(11)"`
	Label      string    `xorm:"not null comment('字典标签') VARCHAR(255)"`
	Value      string    `xorm:"not null comment('字典值') VARCHAR(255)"`
	Sort       string    `xorm:"comment('排序') VARCHAR(255)"`
	DictId     int64     `xorm:"comment('字典id') index BIGINT(11)"`
	CreateTime time.Time `xorm:"comment('创建日期') DATETIME"`
}
