package models

type FilePicture struct {
	Id     int64 `xorm:"pk comment('file_storage表ID') BIGINT(20)"`
	Width  int   `xorm:"not null default 0 comment('宽度') INT(10)"`
	Height int   `xorm:"not null default 0 comment('高度') INT(10)"`
}
