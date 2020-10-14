package lib

import (
	ParentMd5 "crypto/md5"
	ParentSha1 "crypto/sha1"
	ParentSha256 "crypto/sha256"
	ParentSha512 "crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	ParentCrc32 "hash/crc32"
	"strconv"
)

type Crypt struct {
	Str string
}

func (h *Crypt) Md5() string {
	return fmt.Sprintf("%x", ParentMd5.Sum([]byte(h.Str)))
}

func (h *Crypt) Sha1() string {
	return fmt.Sprintf("%x", ParentSha1.Sum([]byte(h.Str)))
}

func (h *Crypt) Sha256() string {
	hash := ParentSha256.New()
	hash.Write([]byte(h.Str))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func (h *Crypt) Sha512() string {
	hash := ParentSha512.New()
	hash.Write([]byte(h.Str))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func (h *Crypt) Crc32() string {
	return strconv.Itoa(int(ParentCrc32.ChecksumIEEE([]byte(h.Str))))
}

func (h *Crypt) Base64Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(h.Str))
}

func (h *Crypt) Base64Decode() string {
	data, err := base64.StdEncoding.DecodeString(h.Str)
	if err != nil {
		return ""
	}
	return string(data)
}

func (h *Crypt) UrlBase64Encode() string {
	return base64.URLEncoding.EncodeToString([]byte(h.Str))
}

func (h *Crypt) UrlBase64Decode() string {
	data, err := base64.URLEncoding.DecodeString(h.Str)
	if err != nil {
		return ""
	}
	return string(data)
}

func (h *Crypt) BiuPwdNewBuilder(password string) (salt string, newPassword string) {
	ConfigLib := new(Config)
	UtilLib := new(Util)
	salt = UtilLib.RandomStr(16)
	h.Str = password
	h.Str = h.Sha512() + ConfigLib.Get("hex_key").String() + salt
	newPassword = h.Sha512()
	return
}

func (h *Crypt) BiuPwdBuilder(salt string, password string) (newPassword string) {
	ConfigLib := new(Config)
	h.Str = password
	h.Str = h.Sha512() + ConfigLib.Get("hex_key").String() + salt
	newPassword = h.Sha512()
	return
}
