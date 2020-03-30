package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Engine *gorm.DB

func init() {
	var err error
	Engine, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/biushop?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Engine.Error != nil {
		fmt.Printf("database error %v", Engine.Error)
	}
}
