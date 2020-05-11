package models

type AuthRoleDept struct {
	RoleId int64 `xorm:"not null pk BIGINT(20)"`
	DeptId int64 `xorm:"not null pk index BIGINT(20)"`
}
