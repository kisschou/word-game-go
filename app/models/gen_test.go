package models

import (
	"time"
)

type GenTest struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"VARCHAR(255)"`
	Sex        int       `xorm:"comment('性别') INT(255)"`
	CreateTime time.Time `xorm:"DATETIME"`
}
