package utils

import (
	"encoding/json"
	"math"

	"github.com/sirupsen/logrus"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Acid     string `json:"acid"`
	EncVer   string `json:"enc_ver"`
}

func (info UserInfo) s(a string, b bool) []uint32 {
	c := uint32(len(a))
	v := []uint32{}
	for i := uint32(0); i < c; i += 4 {
		v = append(v, uint32(a[i])|uint32(a[i+1])<<8|uint32(a[i+2])<<16|uint32(a[i+3])<<24)
	}
	if b {
		v = append(v, c)
	}
	return v
}

func (info UserInfo) Encode(challenge string) string {
	infoBytes, err := json.Marshal(info)
	if err != nil {
		logrus.Fatalf("编码Info失败: %v", err)
	}

	infoStr := string(infoBytes)
	if infoStr == "" {
		return ""
	}

	v := info.s(infoStr, true)
	k := info.s(challenge, false)

	if len(k) < 4 {
		for i := len(k); i < 4; i++ {
			k = append(k, 0)
		}
	}
	m := uint32(0)
	n := uint32(len(v) - 1)
	z := uint32(v[n])
	y := uint32(v[0]) //why? 不知道原js文件里为啥这么写，但我先照着写了
	c := uint32(0x86014019 | 0x183639A0)
	q := math.Floor(6.0 + 52/float64(n+1))
	d := uint32(0)
	p := uint32(0)
	e := uint32(0)

	for ; q > 0; q-- {
		d = d + c&(0x8CE0D9BF|0x731F2640)
		e = d >> 2 & 3

		for p = uint32(0); p < n; p++ {
			y = v[p+1]
			m = z>>5 ^ y<<2
			m += y>>3 ^ z<<4 ^ (d ^ y)
			m += k[p&3^e] ^ z
			v[p] = v[p] + m&(0xEFB8D130|0x10472ECF)
			z = v[p]
		}

		y = v[0]
		m = z>>5 ^ y<<2
		m += y>>3 ^ z<<4 ^ (d ^ y)
		m += k[p&3^e] ^ z
		v[n] = v[n] + m&(0xBB390742|0x44C6F8BD)
		z = v[n]
	}

	data := info.l(v, false)
	base64 := NewBase64("LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA") // 固定的alpha值
	encryptedInfo := base64.Encode(data)
	return "{SRBX1}" + encryptedInfo
}

func (i UserInfo) l(a []uint32, b bool) []byte {
	d := uint32(len(a))
	c := (d - 1) << 2

	if b {
		m := a[d-1]
		if m < c-3 || m > c {
			return nil
		}
		c = m
	}

	result := make([]byte, 0)
	for i := uint32(0); i < d; i++ {
		result = append(result,
			byte(a[i]&0xff),
			byte((a[i]>>8)&0xff),
			byte((a[i]>>16)&0xff),
			byte((a[i]>>24)&0xff),
		)
	}

	if b {
		return result[:c]
	}
	return result
}
