package lib

import (
	ParentMd5 "crypto/md5"
	ParentSha1 "crypto/sha1"
	"encoding/base64"
	"fmt"
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
