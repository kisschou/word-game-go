package models

import (
	"time"
)

type LogVisit struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	CreateTime time.Time `xorm:"DATETIME"`
	Date       string    `xorm:"unique VARCHAR(255)"`
	IpCounts   int64     `xorm:"BIGINT(20)"`
	PvCounts   int64     `xorm:"BIGINT(20)"`
	WeekDay    string    `xorm:"VARCHAR(255)"`
}
