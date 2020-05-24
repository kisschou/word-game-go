package models

import (
	"time"
)

type QiniuContent struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Bucket     string    `xorm:"comment('Bucket 识别符') VARCHAR(255)"`
	Name       string    `xorm:"comment('文件名称') VARCHAR(255)"`
	Size       string    `xorm:"comment('文件大小') VARCHAR(255)"`
	Type       string    `xorm:"comment('文件类型：私有或公开') VARCHAR(255)"`
	UpdateTime time.Time `xorm:"comment('上传或同步的时间') DATETIME"`
	Url        string    `xorm:"comment('文件url') VARCHAR(255)"`
	Suffix     string    `xorm:"VARCHAR(255)"`
}
