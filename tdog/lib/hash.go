package lib

import (
	ParentMd5 "crypto/md5"
	ParentSha1 "crypto/sha1"
	"fmt"
)

type Hash struct {
	Str string
}

func (h *Hash) Md5() string {
	return fmt.Sprintf("%x", ParentMd5.Sum([]byte(h.Str)))
}

func (h *Hash) Sha1() string {
	return fmt.Sprintf("%x", ParentSha1.Sum([]byte(h.Str)))
}
