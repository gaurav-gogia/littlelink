package biz

import (
	"crypto/md5"
	"encoding/hex"
)

func linkHash(link string) string {
	hash := md5.Sum([]byte(link))
	return hex.EncodeToString(hash[:5])
}
