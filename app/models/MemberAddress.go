package models

import (
	_ "github.com/jinzhu/gorm"
	. "wordgame/library/database"
)

type UserAddress struct {
	Id         uint32 `gorm:"primary_key" json:"id" binding:"required"`
	UserId     uint32 `json:"user_id" binding:"required"`
	Receiver   string `json:"receiver" binding:"required"`
	Mobile     string `json:"mobile" binding:"required"`
	ProvinceId uint32 `json:"province_id" binding:"required"`
	CityId     uint32 `json:"city_id" binding:"required"`
	DistrictId uint32 `json:"district_id" binding:"required"`
	Street     string `json:"street" binding:"required"`
	IsDefault  uint32 `json:"is_default" binding:"required"`
	Status     uint32 `json:"status" binding:"required"`
}

func (info *UserAddress) FindAll() (addressList []UserAddress, err error) {
	result := Engine.Find(&addressList)

	if result != nil {
		err = result
		return
	}

	return
}
