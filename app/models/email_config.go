package models

type EmailConfig struct {
	Id       int64  `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	FromUser string `xorm:"comment('收件人') VARCHAR(255)"`
	Host     string `xorm:"comment('邮件服务器SMTP地址') VARCHAR(255)"`
	Pass     string `xorm:"comment('密码') VARCHAR(255)"`
	Port     string `xorm:"comment('端口') VARCHAR(255)"`
	User     string `xorm:"comment('发件者用户名') VARCHAR(255)"`
}
