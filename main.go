package gototp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math"
	"time"
)

const (
	DefaultDigits   = 6
	DefaultInterval = 30
)

type TOTP struct {
	secret   []byte
	digits   int
	interval int64
	format   string
}

func New(s string, d uint, i uint) (*TOTP, error) {
	b, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return &TOTP{
		secret:   b,
		digits:   int(d),
		interval: int64(i),
		format:   fmt.Sprintf("%%0%dd", d),
	}, nil
}

func NewDefault(s string) (*TOTP, error) {
	return New(s, DefaultDigits, DefaultInterval)
}

func (t *TOTP) Generate() string {
	hasher := hmac.New(sha1.New, t.secret)
	hasher.Write(t.time())
	hash := hasher.Sum(nil)

	offset := int(hash[len(hash)-1] & 0xf)
	code := ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)

	code = code % int(math.Pow10(t.digits))
	return fmt.Sprintf(t.format, code)
}

func (t *TOTP) time() []byte {
	n := time.Now().Unix() / t.interval
	b := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n = n >> 8
	}
	return b
}
