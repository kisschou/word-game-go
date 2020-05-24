package models

type AlipayConfig struct {
	Id                   int64  `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	AppId                string `xorm:"comment('应用ID') VARCHAR(255)"`
	Charset              string `xorm:"comment('编码') VARCHAR(255)"`
	Format               string `xorm:"comment('类型 固定格式json') VARCHAR(255)"`
	GatewayUrl           string `xorm:"comment('网关地址') VARCHAR(255)"`
	NotifyUrl            string `xorm:"comment('异步回调') VARCHAR(255)"`
	PrivateKey           string `xorm:"comment('私钥') TEXT"`
	PublicKey            string `xorm:"comment('公钥') TEXT"`
	ReturnUrl            string `xorm:"comment('回调地址') VARCHAR(255)"`
	SignType             string `xorm:"comment('签名方式') VARCHAR(255)"`
	SysServiceProviderId string `xorm:"comment('商户号') VARCHAR(255)"`
}
