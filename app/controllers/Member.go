package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wordgame/app/models"
	. "wordgame/library/cache"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func UAuth(c *gin.Context) {
	header := c.Request.Header

	var member models.User
	err := c.BindJSON(&member)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	memberInfo, err := member.Login("yxbobo", "111111")

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"header":     header,
		"IpAddr":     c.ClientIP(),
		"body":       c.PostForm("username"),
		"memberInfo": memberInfo,
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

	RedisSelectDb(2)
	redisData := Redis.LRange("order:reward:json", 0, 20).Val()

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"memberList":    memberList,
		"memberAddress": memberAddress,
		"redis":         redisData,
		"message":       "success",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Module",
	})
}
