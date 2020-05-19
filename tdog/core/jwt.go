package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"wordgame/tdog/lib"
)

type (
	// 用来表明签名的加密算法 token 类型等
	JwtHeader struct {
		Type      string // 类型
		Algorithm string // 加密算法
	}

	// Payload 记录你需要的信息。 其中应该包含 Claims
	JwtPayload map[string]interface{}

	// 通过 header 生明的加密方法生成 签名
	JwtSignature string

	// jwt 数据
	Jwt struct {
		header    string
		payload   string
		signature string
	}
)

func (header *JwtHeader) New() *JwtHeader {
	return &JwtHeader{
		Type:      "JWT",
		Algorithm: "HS256",
	}
}

/**
 * USAGE:
 * jwt := new(core.Jwt)
 * data := make(map[string]interface{})
 * data["username"] = username
 * data["password"] = password
 * jwt.New(data)
 **/
func (jwt *Jwt) New(data JwtPayload) string {
	crypt := new(lib.Crypt)

	// header
	jwtHeader := make(map[string]string)
	header := new(JwtHeader)
	header = header.New()
	jwtHeader["type"] = header.Type
	jwtHeader["alg"] = header.Algorithm
	jsonData, _ := json.Marshal(jwtHeader)
	crypt.Str = string(jsonData)
	jwt.header = crypt.Base64Encode()

	// payload
	payload := make(map[string]interface{})
	payload["data"] = data
	payload["ita"] = time.Now().Unix()
	payload["exp"] = 7200
	jsonData, _ = json.Marshal(payload)
	crypt.Str = string(jsonData)
	jwt.payload = crypt.Base64Encode()

	// signature
	conf := new(lib.Config)
	jsonData, _ = json.Marshal(payload)
	crypt.Str = string(jsonData) + conf.Get("hex_key").String()
	jwt.signature = crypt.Sha256()

	// save to redis
	redis := new(lib.Redis)
	redis.NewEngine()
	redis.Engine.SetNX("jwt:"+jwt.header+"."+jwt.payload+"."+jwt.signature, 1, time.Duration(7200)*time.Second)

	return jwt.header + "." + jwt.payload + "." + jwt.signature
}

func (jwt *Jwt) Walk(data string) *Jwt {
	jwtData := strings.Split(data, ".")
	return &Jwt{
		header:    jwtData[0],
		payload:   jwtData[1],
		signature: jwtData[2],
	}
}

func (jwt *Jwt) Check(data string) bool {
	// check isExists in redis.
	redis := new(lib.Redis)
	redis.NewEngine()
	if redis.Engine.Exists("jwt:"+data).Val() < 1 {
		return false
	}

	jwt = jwt.Walk(data)
	crypt := new(lib.Crypt)

	// check header.
	header := new(JwtHeader)
	header = header.New()
	crypt.Str = jwt.header
	jwtHeader := make(map[string]string)
	json.Unmarshal([]byte(crypt.Base64Decode()), &jwtHeader)
	if jwtHeader["type"] != header.Type || jwtHeader["alg"] != header.Algorithm {
		return false
	}

	// check payload.
	crypt.Str = jwt.payload
	jwtPayload := make(map[string]interface{})
	json.Unmarshal([]byte(crypt.Base64Decode()), &jwtPayload)
	ita, _ := strconv.Atoi(fmt.Sprintf("%1.0f", jwtPayload["ita"]))
	exp, _ := strconv.Atoi(fmt.Sprintf("%1.0f", jwtPayload["exp"]))
	ita = ita + exp
	if ita < int(time.Now().Unix()) {
		return false
	}

	// check signature.
	conf := new(lib.Config)
	crypt.Str = jwt.payload
	crypt.Str = crypt.Base64Decode() + conf.Get("hex_key").String()
	if jwt.signature != crypt.Sha256() {
		return false
	}

	return true
}

func (jwt *Jwt) Get(data string, key string) interface{} {
	jwt = jwt.Walk(data)
	crypt := new(lib.Crypt)
	crypt.Str = jwt.payload
	jwtPayload := make(map[string]interface{})
	json.Unmarshal([]byte(crypt.Base64Decode()), &jwtPayload)
	list := jwtPayload["data"].(map[string]interface{})
	return list[key]
}
