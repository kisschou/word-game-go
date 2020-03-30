package encrypt

import (
	ParentMd5 "crypto/md5"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	hash := ParentMd5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	return md5str
}

func Sha1(str string) string {
	sha1str := ""
	return sha1str
}
