package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type MySql struct {
}

var SqlEngine *xorm.Engine

func init() {
	var err error

	conf := config.Config{}
	mysqlUser := conf.Get("database.master.user").(string)
	mysqlPass := conf.Get("database.master.pass").(string)
	mysqlHost := conf.Get("database.master.host").(string)
	mysqlPort := conf.Get("database.master.port").(string)
	mysqlDb := conf.Get("database.master.db").(string)
	mysqlCharset := conf.Get("database.master.charset").(string)
	mysqlPrefix := conf.Get("database.master.prefix").(string)

	dsn := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=" + mysqlCharset + "&parseTime=True&loc=Local&timeout=10ms"
	SqlEngine, err = xorm.NewEngine("mysql", dsn)

	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		os.Exit(0)
	}

	if err := SqlEngine.Ping(); err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		os.Exit(0)
	}

	// 日志打印SQL
	SqlEngine.ShowSQL(true)
	// 设置连接池的空闲数大小
	SqlEngine.SetMaxIdleConns(5)
	// 设置最大连接数
	SqlEngine.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	SqlEngine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, mysqlPrefix))
}
