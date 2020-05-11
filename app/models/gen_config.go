package models

type GenConfig struct {
	Id         int64  `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	TableName  string `xorm:"comment('表名') VARCHAR(255)"`
	Author     string `xorm:"comment('作者') VARCHAR(255)"`
	Cover      int    `xorm:"comment('是否覆盖') BIT(1)"`
	ModuleName string `xorm:"comment('模块名称') VARCHAR(255)"`
	Pack       string `xorm:"comment('至于哪个包下') VARCHAR(255)"`
	Path       string `xorm:"comment('前端代码生成的路径') VARCHAR(255)"`
	ApiPath    string `xorm:"comment('前端Api文件路径') VARCHAR(255)"`
	Prefix     string `xorm:"comment('表前缀') VARCHAR(255)"`
	ApiAlias   string `xorm:"comment('接口名称') VARCHAR(255)"`
}
