package utils

type Base64 struct {
	ALPHA string
}

func (b *Base64) Encode(src []byte) string {
	s := string(src)
	x := []byte{}
	imax := len(s) - len(s)%3
	b10 := 0
	if len(s) == 0 {
		return s
	}

	for i := 0; i < imax; i += 3 {
		b10 = (b.getBytes(s, i) << 16) | (b.getBytes(s, i+1) << 8) | (b.getBytes(s, i+2))
		x = append(x, b.ALPHA[b10>>18], b.ALPHA[(b10>>12)&63], b.ALPHA[(b10>>6)&63], b.ALPHA[b10&63])
	}
	switch len(s) - imax {
	case 1:
		b10 = b.getBytes(s, imax) << 16
		x = append(x, b.ALPHA[b10>>18], b.ALPHA[(b10>>12)&63], '=', '=')
	case 2:
		b10 = b.getBytes(s, imax)<<16 | b.getBytes(s, imax+1)<<8
		x = append(x, b.ALPHA[b10>>18], b.ALPHA[(b10>>12)&63], b.ALPHA[(b10>>6)&63], '=')
	}

	return string(x)
}

func (b *Base64) getBytes(s string, i int) int {
	return int(s[i])
}

func NewBase64(alpha string) *Base64 {
	return &Base64{
		ALPHA: alpha,
	}
}
