package models

import (
	"time"
)

type MntDeploy struct {
	Id         int64     `xorm:"pk autoincr comment('编号') BIGINT(20)"`
	AppId      int64     `xorm:"comment('应用编号') index BIGINT(20)"`
	CreateTime time.Time `xorm:"DATETIME"`
}
