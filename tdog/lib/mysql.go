package lib

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type MySql struct {
	Engine      *xorm.Engine
	EngineGroup *xorm.EngineGroup
}

func (mysql *MySql) NewEngine() {
	var err error
	conf := new(Config)

	// master-slave mode.
	if conf.Get("database.master_slave").Bool() {
		mysql.EngineGroup, err = cluster()
	} else {
		mysql.Engine, err = singleton()
	}

	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		os.Exit(0)
	}
}

func cluster() (engineGroup *xorm.EngineGroup, err error) {
	conf := new(Config)
	slaveList := conf.Get("database.slave_list").StringSlice()
	for _, slave := range slaveList {
		fmt.Println("slave: ", slave)
	}

	return
}

func singleton() (engine *xorm.Engine, err error) {
	conf := new(Config)

	mysqlUser := conf.Get("database.master.user").String()
	mysqlPass := conf.Get("database.master.pass").String()
	mysqlHost := conf.Get("database.master.host").String()
	mysqlPort := conf.Get("database.master.port").String()
	mysqlDb := conf.Get("database.master.db").String()
	mysqlCharset := conf.Get("database.master.charset").String()
	mysqlPrefix := ""
	conf.Get("database.master.prefix")
	if conf.IsExists() {
		mysqlPrefix = conf.String()
	}

	dsn := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=" + mysqlCharset + "&parseTime=True&loc=Local&timeout=10ms"
	engine, err = xorm.NewEngine("mysql", dsn)

	if err != nil {
		return
	}

	// 日志打印SQL
	engine.ShowSQL(conf.Get("database.debug").Bool())
	// 设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	// 设置最大连接数
	engine.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, mysqlPrefix))

	return
}

func (mysql *MySql) Ping() bool {
	if err := mysql.Engine.Ping(); err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return false
	}
	return true
}
