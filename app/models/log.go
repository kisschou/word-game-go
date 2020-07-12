package models

import (
	"time"

	"wordgame/tdog/core"
)

type (
	LogModel struct {
		Base core.Model
	}

	Log struct {
		Id              int64     `xorm:"pk autoincr BIGINT(20)"`
		CreateTime      time.Time `xorm:"DATETIME"`
		Description     string    `xorm:"VARCHAR(255)"`
		ExceptionDetail string    `xorm:"TEXT"`
		LogType         string    `xorm:"VARCHAR(255)"`
		Method          string    `xorm:"VARCHAR(255)"`
		Params          string    `xorm:"TEXT"`
		RequestIp       string    `xorm:"VARCHAR(255)"`
		Time            int64     `xorm:"BIGINT(20)"`
		Username        string    `xorm:"VARCHAR(255)"`
		Address         string    `xorm:"VARCHAR(255)"`
		Browser         string    `xorm:"VARCHAR(255)"`
	}
)
