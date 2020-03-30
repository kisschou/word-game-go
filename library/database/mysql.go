package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/biushop?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := Engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	// 日志打印SQL
	Engine.ShowSQL(true)
	// 设置连接池的空闲数大小
	Engine.SetMaxIdleConns(5)
	// 设置最大打开连接数
	Engine.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	Engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "biu_"))
}
