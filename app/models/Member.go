package models

import (
	_ "github.com/jinzhu/gorm"
	. "wordgame/library/database"
	"wordgame/library/encrypt"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Salt     string `json:"salt" binding:"required"`
}

func (info *User) Insert() (userId uint, err error) {
	result, err := Engine.Insert(&info)
	if result != 0 {
		return
	}
	return
}

func (info *User) FindAll() (memberList []User, err error) {
	result := Engine.Find(&memberList)

	if result != nil {
		err = result
		return
	}

	return
}

func thinkUcenterMd5(str string, key string, salt string) string {
	password := encrypt.Md5(str)
	return password
}
