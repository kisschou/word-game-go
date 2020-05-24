package models

type Shop struct {
	Id            int    `xorm:"not null pk autoincr INT(11)"`
	CreateTime    int    `xorm:"not null INT(11)"`
	Introduction  string `xorm:"not null VARCHAR(255)"`
	IsJd          int    `xorm:"not null INT(11)"`
	JdAccessToken string `xorm:"not null VARCHAR(255)"`
	JdAccount     string `xorm:"not null VARCHAR(255)"`
	MerchantId    int    `xorm:"not null INT(11)"`
	Status        int    `xorm:"not null INT(11)"`
	Title         string `xorm:"not null VARCHAR(255)"`
	UpdateTime    int    `xorm:"not null INT(11)"`
}
