package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Level int
	Key   string
}

func buildFilePath(logger *Logger) (filePath string) {
	var util *Util
	levelMap := map[int]string{0: "logs", 1: "runtime"}
	filePath, _ = os.Getwd()
	crypt := Crypt{Str: logger.Key}
	filePath += "/" + levelMap[logger.Level] + "/" + crypt.Md5() + "/"
	if logger.Level == 0 {
		filePath += time.Now().Format("2006-01-02") + "/" //  2006-01-02 15:04:05
	}
	util.DirExistsAndCreate(filePath)
	return
}

func buildFileName(logger *Logger) (fileName string) {
	crypt := Crypt{Str: logger.Key}
	fileName = crypt.Sha1() + ".log"
	return
}

func toFile(logger *Logger, context string) {
	logFilePath := buildFilePath(logger)
	logFileName := buildFileName(logger)

	fileName := logFilePath + logFileName
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("write to log file FAIL: ", err)
		os.Exit(0)
	}

	// 实例化
	log := logrus.New()
	// 设置输出
	log.Out = src
	/**
	 * 日志的级别（来自@dylanbeattie）
	 * - Fatal：网站挂了，或者极度不正常
	 * - Error：跟遇到的用户说对不起，可能有bug
	 * - Warn：记录一下，某事又发生了
	 * - Info：提示一切正常
	 * - debug：没问题，就看看堆栈
	 */
	log.SetLevel(logrus.InfoLevel)
	// 设置输出日志中添加文件名和方法信息
	log.SetReportCaller(true)
	// 设置日志格式
	log.SetFormatter(&logrus.TextFormatter{})
	// 写入日志
	log.Infof("%s", context)
}

func (logger *Logger) New(context string) {
	toFile(logger, context)
}
