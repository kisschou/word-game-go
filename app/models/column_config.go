package models

type ColumnConfig struct {
	Id             int64  `xorm:"pk autoincr BIGINT(20)"`
	TableName      string `xorm:"VARCHAR(255)"`
	ColumnName     string `xorm:"VARCHAR(255)"`
	ColumnType     string `xorm:"VARCHAR(255)"`
	DictName       string `xorm:"VARCHAR(255)"`
	Extra          string `xorm:"VARCHAR(255)"`
	FormShow       int    `xorm:"BIT(1)"`
	FormType       string `xorm:"VARCHAR(255)"`
	KeyType        string `xorm:"VARCHAR(255)"`
	ListShow       int    `xorm:"BIT(1)"`
	NotNull        int    `xorm:"BIT(1)"`
	QueryType      string `xorm:"VARCHAR(255)"`
	Remark         string `xorm:"VARCHAR(255)"`
	DateAnnotation string `xorm:"VARCHAR(255)"`
}
