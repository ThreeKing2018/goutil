package strtool

import (
	"fmt"
	"testing"
)

func TestTrimRightSpace(t *testing.T) {
	str := "aaabc\n\t\r"
	rs := TrimRightSpace(str)
	fmt.Println(rs)
}
func TestMd5(t *testing.T) {
	str := RandomString(10)
	rs := Md5(str)
	fmt.Println(rs)
}
func TestRandomString(t *testing.T) {
	str := RandomString(5)
	fmt.Println(str)
}
