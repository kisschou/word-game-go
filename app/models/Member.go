package models

import (
	_ "github.com/jinzhu/gorm"
	. "wordgame/library/database"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}

func (User) TableName() string {
	return "biu_user"
}

func (info *User) Insert() (userId uint, err error) {
	result := Engine.Create(&info)
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (info *User) FindAll() (memberList []User, err error) {
	result := Engine.Find(&memberList)

	if result.Error != nil {
		err = result.Error
		return
	}

	return
}
