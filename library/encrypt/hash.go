package encrypt

import (
	ParentMd5 "crypto/md5"
	ParentSha1 "crypto/sha1"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	md5str := fmt.Sprintf("%x", ParentMd5.Sum(data))
	return md5str
}

func Sha1(str string) string {
	data := []byte(str)
	sha1str := fmt.Sprintf("%x", ParentSha1.Sum(data))
	return sha1str
}

func ThinkUcenterMd5(str string, key string, salt string) string {
	str = Sha1(str)
	str += key
	str += salt
	return Md5(str)
}
