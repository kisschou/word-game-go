package models

import (
	"time"
)

type QuartzLog struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	BaenName        string    `xorm:"VARCHAR(255)"`
	CreateTime      time.Time `xorm:"DATETIME"`
	CronExpression  string    `xorm:"VARCHAR(255)"`
	ExceptionDetail string    `xorm:"TEXT"`
	IsSuccess       int       `xorm:"BIT(1)"`
	JobName         string    `xorm:"VARCHAR(255)"`
	MethodName      string    `xorm:"VARCHAR(255)"`
	Params          string    `xorm:"VARCHAR(255)"`
	Time            int64     `xorm:"BIGINT(20)"`
}
