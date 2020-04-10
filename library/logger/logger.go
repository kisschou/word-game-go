package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	. "wordgame/library/encrypt"
	. "wordgame/library/unit"
)

func LoggerAccess() gin.HandlerFunc {
	logFilePath, _ := os.Getwd()
	logFilePath += "/logs/"
	logFileName := "access.log"

	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func buildFilePath(key string, level int) (filePath string) {
	levelMap := map[int]string{0: "logs", 1: "runtime"}

	filePath, _ = os.Getwd()
	filePath += "/" + levelMap[level] + "/" + Md5(key) + "/"
	if level == 0 {
		filePath += time.Now().Format("2006-01-02") + "/" // 2006-01-02 15:04:05
	}

	DirExistsAndCreate(filePath)
	return
}

func buildFileName(key string) (fileName string) {
	fileName = Sha1(key)
	fileName += ".log"
	return
}

func LoggerToFile(key string, context string, level int) {
	logFilePath := buildFilePath(key, level)
	logFileName := buildFileName(key)

	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	// 写入日志
	logger.Infof("%s", context)
}

func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
