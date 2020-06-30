package lib

import (
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// 生成指定数量随机字母加数字
func (u *Util) RandomStr(length int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 判断map中是否存在某个key
func (u *Util) InMap(key string, dataMap map[string]interface{}) bool {
	if _, ok := dataMap[key]; ok {
		return true
	}
	return false
}

func (u *Util) InStringSlice(key string, dataStringSlice []string) bool {
	for _, v := range dataStringSlice {
		if v == key {
			return true
		}
	}
	return false
}

// 检测字符串是邮件、手机号、字符串
// 返回: 0字符串1邮件2手机号
func (u *Util) CheckStrType(str string) int {
	if verifyEmailFormat(str) {
		return 1
	}

	if verifyPhoneFormat(str) {
		return 2
	}

	return 0
}

func verifyEmailFormat(email string) bool {
	// pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func verifyPhoneFormat(phone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

// 获取设备id
// 通过网卡ipv4生成
func (u *Util) GetMachineId() int64 {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 127000000001
	}

	gatewayIp := ""
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				gatewayIp = ipnet.IP.String()
				break
			}
		}
	}

	gatewayIpSplit := strings.Split(gatewayIp, ".")
	gatewayIp = ""
	for _, v := range gatewayIpSplit {
		if len(v) < 3 {
			for i := len(v); i < 3; i++ {
				v = "0" + v
			}
		}
		gatewayIp += v
	}
	machineId, _ := strconv.ParseInt(gatewayIp, 10, 64)
	return machineId
}

func (u *Util) RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
