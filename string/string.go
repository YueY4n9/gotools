package _string

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5(str string) string {
	m := md5.New()
	_, _ = io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Sha256(str string) string {
	if str == "" {
		return str
	}
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
