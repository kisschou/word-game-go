package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wordgame/app/models"
	. "wordgame/library/encrypt"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func Login(c *gin.Context) {
	var member models.User     // 定义json 变量 数据结构类型 为 (models/member).MemberInfo
	err := c.BindJSON(&member) // 获取前台传过来的json数据

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	memberList, err := member.FindAll()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", memberList)
		return
	}

	var address models.UserAddress
	err = c.BindJSON(&address)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	memberAddress, err := address.FindAll()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", memberList)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"md5":           Md5("123456"),
		"memberList":    memberList,
		"memberAddress": memberAddress,
		"message":       "success",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Module",
	})
}
