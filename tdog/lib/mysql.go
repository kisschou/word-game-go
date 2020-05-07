package lib

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type (
	MySql struct {
		IsCluster   bool
		Engine      *xorm.Engine
		EngineGroup *xorm.EngineGroup
	}

	MySqlConf struct {
		engine  string
		host    string
		port    string
		user    string
		pass    string
		db      string
		charset string
		prefix  string
		dsn     string
		debug   bool
	}
)

func (mysql *MySql) NewEngine() {
	var err error
	conf := new(Config)

	// master-slave mode.
	mysql.IsCluster = conf.Get("database.master_slave").Bool()
	if mysql.IsCluster {
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
	conns := []string{}
	// append master conf.
	var myCnf *MySqlConf
	myCnf = myCnf.Get("master")
	conns = append(conns, myCnf.dsn)

	// append slave conf.
	conf := new(Config)
	slaveList := conf.Get("database.slave_list").StringSlice()
	for _, slave := range slaveList {
		myCnf = myCnf.Get(slave)
		conns = append(conns, myCnf.dsn)
	}

	/**
	 * 负载策略
	 *
	 * 1. 随机访问负载策略 xorm.RandomPolicy()
	 * 2. 权重随机访问负载策略 xorm.WeightRandomPolicy([]int{2, 3})
	 * 3. 轮询访问负载策略 xorm.RoundRobinPolicy()
	 * 4. 权重轮询访问负载策略 xorm.WeightRoundRobinPolicy([]int{2, 3})
	 * 5. 最小连接数访问负载策略 xorm.LeastConnPolicy()
	 * 6. 自定义
	 *		type GroupPolicy interface {
	 *			Slave(*EngineGroup) *Engine
	 *		}
	 *
	 * 注: 当前采用随机访问负载策略
	 */
	engineGroup, err = xorm.NewEngineGroup(myCnf.engine, conns, xorm.RandomPolicy())

	// 日志打印SQL
	engineGroup.ShowSQL(myCnf.debug)
	// 设置连接池的空闲数大小
	engineGroup.SetMaxIdleConns(5)
	// 设置最大连接数
	engineGroup.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engineGroup.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, myCnf.prefix))

	return
}

func singleton() (engine *xorm.Engine, err error) {
	var myCnf *MySqlConf
	myCnf = myCnf.Get("master")
	engine, err = xorm.NewEngine(myCnf.engine, myCnf.dsn)

	if err != nil {
		return
	}

	// 日志打印SQL
	engine.ShowSQL(myCnf.debug)
	// 设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	// 设置最大连接数
	engine.SetMaxOpenConns(5)
	// 名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, myCnf.prefix))

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

func (myCnf *MySqlConf) Get(hostType string) *MySqlConf {
	conf := new(Config)
	mysqlUser := conf.Get("database.master.user").String()
	mysqlPass := conf.Get("database." + hostType + ".pass").String()
	mysqlHost := conf.Get("database." + hostType + ".host").String()
	mysqlPort := conf.Get("database." + hostType + ".port").String()
	mysqlDb := conf.Get("database." + hostType + ".db").String()
	mysqlCharset := conf.Get("database." + hostType + ".charset").String()
	mysqlPrefix := ""
	conf.Get("database." + hostType + ".prefix")
	if conf.IsExists() {
		mysqlPrefix = conf.String()
	}
	dsn := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=" + mysqlCharset + "&parseTime=True&loc=Local&timeout=10ms"
	debug := conf.Get("database.debug").Bool()

	return &MySqlConf{
		engine:  "mysql",
		host:    mysqlHost,
		port:    mysqlPort,
		user:    mysqlUser,
		pass:    mysqlPass,
		db:      mysqlDb,
		charset: mysqlCharset,
		prefix:  mysqlPrefix,
		dsn:     dsn,
		debug:   debug,
	}
}
