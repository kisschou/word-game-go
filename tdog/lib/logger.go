package lib

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Level int
	Key   string
}

func buildFilePath(logger *Logger) (filePath string) {
	levelMap := map[int]string{0: "logs", 1: "runtime"}
	filePath, _ = os.Getwd()
	hash := Hash{Str: logger.Key}
	filePath += "/" + levelMap[logger.Level] + "/" + hash.Md5() + "/"
	if level == 0 {
		filePath += time.Now().Format("2006-01-02") + "/" //  2006-01-02 15:04:05
	}
	return
}

func buildFileName(logger *Logger) (fileName string) {
	hash := Hash{Str: logger.Key}
	fileName = hash.Sha1() + ".log"
}

func toFile(logger *Logger, context string) {
	logFilePath := buildFilePath(logger)
	logFileName := buildFileName(logger)

	fileName := strings.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("write to log file FAIL: ", err)
		os.Exit(0)
	}

	// 实例化
	log := logrus.New()
	// 设置输出
	log.Out = src
	// 设置日志级别
	log.SetLevel(logrus.InfoLevel)
	// 设置日志格式
	log.SetFormatter(&logrus.TextFormatter{})
	// 写入日志
	log.Infof("%s", context)
}

func (logger *Logger) New(context string) {
	toFile(logger, context)
}
