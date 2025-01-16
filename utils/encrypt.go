package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func EncryptPassword(challenge, password string) string {
	// 使用HMAC-MD5算法对密码进行加密
	h := hmac.New(md5.New, []byte(challenge))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
