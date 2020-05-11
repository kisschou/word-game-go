package models

import (
	"time"
)

type AuthMenu struct {
	Id            int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	IFrame        int       `xorm:"comment('是否外链') BIT(1)"`
	Name          string    `xorm:"comment('菜单名称') VARCHAR(255)"`
	Component     string    `xorm:"comment('组件') VARCHAR(255)"`
	Pid           int64     `xorm:"not null comment('上级菜单ID') index BIGINT(20)"`
	Sort          int64     `xorm:"comment('排序') BIGINT(20)"`
	Icon          string    `xorm:"comment('图标') VARCHAR(255)"`
	Path          string    `xorm:"comment('链接地址') VARCHAR(255)"`
	Cache         int       `xorm:"default b'0' comment('缓存') BIT(1)"`
	Hidden        int       `xorm:"default b'0' comment('隐藏') BIT(1)"`
	ComponentName string    `xorm:"default '-' comment('组件名称') VARCHAR(20)"`
	CreateTime    time.Time `xorm:"comment('创建日期') DATETIME"`
	Permission    string    `xorm:"comment('权限') VARCHAR(255)"`
	Type          int       `xorm:"comment('类型') INT(11)"`
}
