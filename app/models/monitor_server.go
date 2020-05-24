package models

type MonitorServer struct {
	Id        int     `xorm:"not null pk autoincr INT(11)"`
	CpuCore   int     `xorm:"comment('CPU内核数') INT(11)"`
	CpuRate   float64 `xorm:"comment('CPU使用率') DOUBLE"`
	DiskTotal float64 `xorm:"comment('磁盘总量') DOUBLE"`
	DiskUsed  float64 `xorm:"comment('磁盘使用量') DOUBLE"`
	MemTotal  float64 `xorm:"comment('内存总数') DOUBLE"`
	MemUsed   float64 `xorm:"comment('内存使用量') DOUBLE"`
	Name      string  `xorm:"comment('名称') VARCHAR(255)"`
	Port      int     `xorm:"comment('访问端口') INT(11)"`
	Sort      int     `xorm:"comment('排序') INT(11)"`
	State     string  `xorm:"comment('状态') VARCHAR(255)"`
	SwapTotal float64 `xorm:"comment('交换区总量') DOUBLE"`
	SwapUsed  float64 `xorm:"comment('交换区使用量') DOUBLE"`
	Address   string  `xorm:"not null comment('服务地址') VARCHAR(255)"`
}
