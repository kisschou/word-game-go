package models

type QiniuConfig struct {
	Id        int64  `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	AccessKey string `xorm:"comment('accessKey') TEXT"`
	Bucket    string `xorm:"comment('Bucket 识别符') VARCHAR(255)"`
	Host      string `xorm:"not null comment('外链域名') VARCHAR(255)"`
	SecretKey string `xorm:"comment('secretKey') TEXT"`
	Type      string `xorm:"comment('空间类型') VARCHAR(255)"`
	Zone      string `xorm:"comment('机房') VARCHAR(255)"`
}
