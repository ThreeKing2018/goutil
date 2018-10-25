package filetool

import (
	"io/ioutil"
	"strings"
)

// ReadFileToBytes reads data type '[]byte' from file by given path.
// It returns error when fail to finish operation.
func ReadFileToBytes(filePath string) ([]byte, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

// ReadFileToString reads data type 'string' from file by given path.
// It returns error when fail to finish operation.
func ReadFileToString(filePath string) (string, error) {
	b, err := ReadFileToBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//读取文件里的数据,去掉空格,换行,制表符,回车
func ReadFileToStringNoLn(filePath string) (string, error) {
	str, err := ReadFileToString(filePath)
	if err != nil {
		return "", err
	}
	str = strings.Trim(string(str), " ")
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str, nil
}
