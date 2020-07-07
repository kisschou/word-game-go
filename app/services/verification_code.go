package services

import (
	"encoding/json"
	"errors"
	"time"

	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type (
	VerificationCode struct {
		Base   core.Service
		IpAddr string
	}

	VerificationCodeKey struct {
		Service              string // 服务标识
		VerificationCodeType string // 验证码类型(短信、邮件、其他)
		Extras               string // 附加数据(根据验证码类型填写手机、邮箱、随机码)
		IpAddr               string // 请求方IP地址
		CreateTime           int64  // 构建时间戳
	}
)

func (verificationCodeKey *VerificationCodeKey) BuildKey() string {
	UtilLib := new(lib.Util)
	captchaAccess := ""
	captchaAccessParams := UtilLib.StructToMap(verificationCodeKey)
	codeKeyByte, _ := json.Marshal(captchaAccessParams)
	CryptLib := new(lib.Crypt)
	CryptLib.Str = string(codeKeyByte)
	captchaAccess = CryptLib.Sha1()
	return captchaAccess
}

func (verificationCode *VerificationCode) Build(service, verificationCodeType, extras, ipAddr string) string {
	verificationCodeKey := new(VerificationCodeKey)
	verificationCodeKey.Service = service
	verificationCodeKey.VerificationCodeType = verificationCodeType
	verificationCodeKey.Extras = extras
	verificationCodeKey.IpAddr = ipAddr
	verificationCodeKey.CreateTime = time.Now().UnixNano()
	captchaAccess := verificationCodeKey.BuildKey()

	UtilLib := new(lib.Util)
	code := UtilLib.RandomNum(4)

	// save to redis
	ConfigLib := new(lib.Config)
	captchaPrefix := ConfigLib.Get("captcha.prefix").String()
	captchaLiveTime := ConfigLib.Get("captcha.live_time").Int()
	key := "captcha:"
	if len(captchaPrefix) > 0 && captchaPrefix != ":" {
		key += captchaPrefix + ":"
	}
	key += captchaAccess
	if captchaLiveTime < 30 {
		captchaLiveTime = 180
	}
	verificationCode.Base.Redis.NewEngine()
	verificationCode.Base.Redis.Engine.SetNX(key, code, time.Duration(captchaLiveTime)*time.Second)

	return captchaAccess
}

func (verificationCode *VerificationCode) GetCode(captchaAccess string) (code string, err error) {
	ConfigLib := new(lib.Config)
	captchaPrefix := ConfigLib.Get("captcha.prefix").String()
	key := "captcha:"
	if len(captchaPrefix) > 0 && captchaPrefix != ":" {
		key += captchaPrefix + ":"
	}
	key += captchaAccess
	verificationCode.Base.Redis.NewEngine()
	if verificationCode.Base.Redis.Engine.Exists(key).Val() < 1 {
		err = errors.New("ERROR_CAPTCHA_TIMEOUT")
		return
	}
	code = verificationCode.Base.Redis.Engine.Get(key).Val()
	return
}

func (verificationCode *VerificationCode) Verify(captchaAccess, code string) (err error) {
	ConfigLib := new(lib.Config)
	captchaPrefix := ConfigLib.Get("captcha.prefix").String()
	key := "captcha:"
	if len(captchaPrefix) > 0 && captchaPrefix != ":" {
		key += captchaPrefix + ":"
	}
	key += captchaAccess
	verificationCode.Base.Redis.NewEngine()
	if verificationCode.Base.Redis.Engine.Exists(key).Val() < 1 {
		err = errors.New("ERROR_CAPTCHA_TIMEOUT")
		return
	}
	existsCode := verificationCode.Base.Redis.Engine.Get(key).Val()
	if code != existsCode {
		err = errors.New("ERROR_CAPTCHA_FAIL")
		return
	}
	return
}
