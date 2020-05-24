package models

import (
	"time"
)

type QuartzJob struct {
	Id             int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	BeanName       string    `xorm:"comment('Spring Bean名称') VARCHAR(255)"`
	CronExpression string    `xorm:"comment('cron 表达式') VARCHAR(255)"`
	IsPause        int       `xorm:"comment('状态：1暂停、0启用') BIT(1)"`
	JobName        string    `xorm:"comment('任务名称') VARCHAR(255)"`
	MethodName     string    `xorm:"comment('方法名称') VARCHAR(255)"`
	Params         string    `xorm:"comment('参数') VARCHAR(255)"`
	Remark         string    `xorm:"comment('备注') VARCHAR(255)"`
	CreateTime     time.Time `xorm:"comment('创建日期') DATETIME"`
}
