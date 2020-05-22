package lib

import (
	"os"
)

type Util struct {
}

// 判断文件是否存在
func (u *Util) FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断是否是目录
func (u *Util) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断是否是文件
func (u *Util) IsFile(path string) bool {
	return !u.IsDir(path)
}

// 判断文件夹是否存在,不存在就创建
func (u *Util) DirExistsAndCreate(path string) {
	if !u.FileExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger := Logger{Level: 0, Key: "error"}
			logger.New(err.Error())
			os.Exit(0)
		}
	}
}
