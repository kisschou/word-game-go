package models

type ShopProduct struct {
	Id                   int64  `xorm:"pk autoincr BIGINT(20)"`
	ShopId               int    `xorm:"not null default 0 comment('店铺ID') INT(10)"`
	CategoryId           int    `xorm:"not null default 0 comment('分类ID') INT(10)"`
	CoverId              int    `xorm:"not null default 0 comment('封面图ID') INT(10)"`
	BrandId              int    `xorm:"not null default 0 comment('品牌ID') INT(10)"`
	Stock                int    `xorm:"not null default 0 comment('库存数量') INT(10)"`
	Sold                 int    `xorm:"not null default 0 comment('卖出数量') INT(10)"`
	SerialNumber         string `xorm:"not null default '' comment('商品编码') VARCHAR(64)"`
	ItemNumber           string `xorm:"not null default '' comment('商品货号') VARCHAR(64)"`
	BarCode              string `xorm:"not null default '' comment('商品条形码') VARCHAR(64)"`
	Title                string `xorm:"not null default '' comment('商品标题') VARCHAR(64)"`
	Price                string `xorm:"not null default 0.00 comment('价格') DECIMAL(10,2)"`
	CostPrice            string `xorm:"not null default 0.00 comment('成本价') DECIMAL(10,2)"`
	MarketPrice          string `xorm:"not null default 0.00 comment('市场价') DECIMAL(10,2)"`
	About                string `xorm:"not null default '' comment('商品简介') VARCHAR(64)"`
	Description          string `xorm:"not null comment('描述') TEXT"`
	MobileDescription    string `xorm:"not null comment('手机端描述') TEXT"`
	AfterSaleDescription string `xorm:"not null comment('售后') TEXT"`
	Weight               string `xorm:"not null default 0.00 comment('重量（KG）') DECIMAL(10,2)"`
	Length               int    `xorm:"not null default 0 comment('长度（MM）') INT(10)"`
	Width                int    `xorm:"not null default 0 comment('宽度（MM）') INT(10)"`
	Height               int    `xorm:"not null default 0 comment('高度（MM）') INT(10)"`
	DeliveryId           int    `xorm:"not null default 0 comment('运费模版ID') INT(10)"`
	IsJd                 int    `xorm:"not null default 0 comment('是否京东商品') TINYINT(1)"`
	JdId                 string `xorm:"not null default '' comment('京东内部商品ID') CHAR(12)"`
	CreateTime           int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime           int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status               int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
}
