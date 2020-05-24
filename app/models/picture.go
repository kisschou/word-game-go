package models

import (
	"time"
)

type Picture struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	CreateTime time.Time `xorm:"comment('上传日期') DATETIME"`
	DeleteUrl  string    `xorm:"comment('删除的URL') VARCHAR(255)"`
	Filename   string    `xorm:"comment('图片名称') VARCHAR(255)"`
	Height     string    `xorm:"comment('图片高度') VARCHAR(255)"`
	Size       string    `xorm:"comment('图片大小') VARCHAR(255)"`
	Url        string    `xorm:"comment('图片地址') VARCHAR(255)"`
	Username   string    `xorm:"comment('用户名称') VARCHAR(255)"`
	Width      string    `xorm:"comment('图片宽度') VARCHAR(255)"`
	Md5code    string    `xorm:"comment('文件的MD5值') VARCHAR(255)"`
}
