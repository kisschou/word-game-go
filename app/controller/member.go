package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Member struct {
}

func (member *Member) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func (member *Member) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login success!!!",
	})
}
