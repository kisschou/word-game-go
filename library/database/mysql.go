package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"wordgame/library/config"
	"wordgame/library/logger"
)

var Engine *xorm.Engine

func init() {
	var err error

	c := config.Config{}
	mysqlUser := c.Get("database.mysql.user").(string)
	mysqlPass := c.Get("database.mysql.pass").(string)
	mysqlHost := c.Get("database.mysql.host").(string)
	mysqlPort := c.Get("database.mysql.port").(string)
	mysqlDb := c.Get("database.mysql.db").(string)
	mysqlCharset := c.Get("database.mysql.charset").(string)
	mysqlPrefix := c.Get("database.mysql.prefix").(string)

	dsn := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=" + mysqlCharset + "&parseTime=True&loc=Local&timeout=10ms"
	Engine, err = xorm.NewEngine("mysql", dsn)

	if err != nil {
		logger.LoggerToFile("error", err.Error(), 0)
		return
	}

	if err := Engine.Ping(); err != nil {
		logger.LoggerToFile("error", err.Error(), 0)
		return
	}

	// 日志打印SQL
	Engine.ShowSQL(true)
	// 设置连接池的空闲数大小
	Engine.SetMaxIdleConns(5)
	// 设置最大打开连接数
	Engine.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	Engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, mysqlPrefix))
}
