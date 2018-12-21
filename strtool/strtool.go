package strtool

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//去掉换行,空格,回车,制表
func TrimRightSpace(s string) string {
	return strings.TrimRight(string(s), "\r\n\t ")
}

//生成MD5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//随机字符串
func RandomString(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
