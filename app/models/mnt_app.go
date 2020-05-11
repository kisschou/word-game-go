package models

import (
	"time"
)

type MntApp struct {
	Id           int64     `xorm:"pk autoincr comment('应用ID') BIGINT(20)"`
	Name         string    `xorm:"comment('应用名称') VARCHAR(255)"`
	UploadPath   string    `xorm:"comment('上传目录') VARCHAR(255)"`
	DeployPath   string    `xorm:"comment('部署路径') VARCHAR(255)"`
	BackupPath   string    `xorm:"comment('备份路径') VARCHAR(255)"`
	Port         int       `xorm:"comment('应用端口') INT(255)"`
	StartScript  string    `xorm:"comment('启动脚本') VARCHAR(4000)"`
	DeployScript string    `xorm:"comment('部署脚本') VARCHAR(4000)"`
	CreateTime   time.Time `xorm:"comment('创建日期') DATETIME"`
}
