package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/Mmx233/BitSrunLoginGo/pkg/srun"
	"github.com/sirupsen/logrus"
)

func EncryptPassword(challenge, password string) string {
	// 使用HMAC-MD5算法对密码进行加密
	h := hmac.New(md5.New, []byte(challenge))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeUserInfo(info any, challenge string) string {
	// 将info编码为JSON格式，并使用XEncode进行加密
	infoBytes, err := json.Marshal(info)
	if err != nil {
		logrus.Fatalf("编码Info失败: %v", err)
	}
	encryptedInfo := srun.Base64(srun.XEncode(string(infoBytes), challenge))
	return "{SRBX1}" + encryptedInfo
}
