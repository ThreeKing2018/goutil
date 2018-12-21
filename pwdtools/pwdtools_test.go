package pwdtools

import (
	"fmt"
	"testing"
)

//获取执行文件目录, 直接会显示出临时文件 ,必须先go build 再go run main.go
func TestGetCurrentDirectory(t *testing.T) {
	fmt.Println(GetCurrentDirectory())
}
func TestGetExecFilePath(t *testing.T) {
	fmt.Println(GetExecFilePath())
}
func TestGetRootDir(t *testing.T) {
	fmt.Println(GetRootDir())
}
