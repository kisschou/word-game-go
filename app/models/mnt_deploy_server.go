package models

type MntDeployServer struct {
	DeployId int64 `xorm:"not null pk BIGINT(20)"`
	ServerId int64 `xorm:"not null pk index BIGINT(20)"`
}
