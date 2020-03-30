package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"wordgame/app/models"
)

type MemberInfo struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func Login(c *gin.Context) {
	var json models.User     // 定义json 变量 数据结构类型 为 (models/member).MemberInfo
	err := c.BindJSON(&json) // 获取前台传过来的json数据

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	memberList, err := json.FindAll()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", memberList)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"memberList": memberList,
		"message":    "success",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Module",
	})
}
