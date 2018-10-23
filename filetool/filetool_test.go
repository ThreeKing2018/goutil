package filetool

import (
	"testing"
	"fmt"
)

func TestFileMTime(t *testing.T) {
	filePath := "test.txt"
	rs , err := FileToInt64(filePath)
	if err != nil {
		t.Error("FileToInt64 err", err.Error())
	}
	fmt.Println(rs)
}
