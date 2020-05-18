package core

import (
	"encoding/json"
	"fmt"
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

func (header *JwtHeader) New() {
}

func (jwt *Jwt) New(data JwtPayload) {
	payload := make(map[string]interface{})
	payload["data"] = data
	payload["ita"] = time.Now().Unix()
	payload["exp"] = 7200

	crypt := new(lib.Crypt)
	jsonData, _ := json.Marshal(payload)
	crypt.Str = string(jsonData)
	fmt.Println(crypt.Base64Encode())
}

func (jwt *Jwt) Get(data string) {
}
