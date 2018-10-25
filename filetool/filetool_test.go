package filetool

import (
	"testing"
	"fmt"
	"io/ioutil"
	"code.sgfoot.com/goutil/datetime"
)

var filePath = "/tmp/test.txt"
func init() {
	content := "1001"
	data := []byte(content)
	if ioutil.WriteFile(filePath, data, 0666) == nil {
		fmt.Println("写入成功")
	}
}
//将文件内的数字转换成int64
func TestFileToInt64(t *testing.T) {
	rs , err := FileToInt64(filePath)
	if err != nil {
		t.Error("FileToInt64 err", err.Error())
	}
	fmt.Println(rs)
}
//将文件内的数字转换成无符号的int64
func TestFileToUint64(t *testing.T) {
	rs , err := FileToUint64(filePath)
	if err != nil {
		t.Error("FileToUint64 err", err.Error())
	}
	fmt.Println(rs)
}
//获取文件名称
func TestBasename(t *testing.T) {
	rs := Basename(filePath)
	fmt.Println(rs)
}
//获取文件修改时间戳
func TestFileMTime(t *testing.T) {
	rs, err := FileMTime(filePath)
	if err != nil  {
		t.Errorf("FileMTime err: %v \n", err.Error())
	}
	fmt.Println(datetime.GetDate(rs))
}
//获取当前目录
func TestDir(t *testing.T) {
	rs := Dir(filePath)
	fmt.Println(rs)
}
func TestDirsUnder(t *testing.T) {
	rs, err := DirsUnder(filePath)
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}