package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wordgame/app/models"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func UAuth(c *gin.Context) {
	var member models.User
	err := member.UAuth(c.PostForm("token"))

	if err != nil {
		fmt.Println(err)
		c.JSON(310, gin.H{"message": string(err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"header":  c.Request.Header,
		"IpAddr":  c.ClientIP(),
		"body":    c.PostForm("token"),
	})
}

func GetInfo(c *gin.Context) {
	var member models.User // 定义json 变量 数据结构类型 为 (models/member).MemberInfo
	member.OpenId = c.PostForm("openId")
	memberInfo, err := member.GetInfo()
	if err != nil {
		c.JSON(310, gin.H{"message": string(err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    memberInfo,
		"message": "success",
	})
}

func Login(c *gin.Context) {
	var member models.User // 定义json 变量 数据结构类型 为 (models/member).MemberInfo

	member.Username = c.PostForm("username")
	member.Password = c.PostForm("password")

	memberInfo, err := member.Login()

	if err != nil {
		c.JSON(310, gin.H{"message": string(err.Error())})
		return
	}

	var address models.UserAddress
	err = c.BindJSON(&address)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    memberInfo,
		"message": "success",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Module",
	})
}
