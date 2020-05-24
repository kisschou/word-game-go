package models

type AuthRoleMenu struct {
	MenuId int64 `xorm:"not null pk comment('菜单ID') BIGINT(20)"`
	RoleId int64 `xorm:"not null pk comment('角色ID') index BIGINT(20)"`
}
