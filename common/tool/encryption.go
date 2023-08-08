package tool

import (
	"crypto/md5"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// MD5加密
func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

// bcrypt加密
func BcryptByString(str string) string {
	byteArr := []byte(str)
	b, err := bcrypt.GenerateFromPassword(byteArr, 4)
	if err != nil {
		panic(err)
	}
	return string(b)
}
