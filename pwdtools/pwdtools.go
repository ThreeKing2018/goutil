package pwdtools

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
	"os/exec"
	"github.com/ThreeKing2018/goutil/logtool"
)

//获取当前执行文件的目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//返回绝对路径
	filepath.Dir(os.Args[0])
	if err != nil {
		logtool.Fetal(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
//获取执行文件当前的目录 最后带/
func GetRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}
//获取执行文件的路径
func GetExecFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file, err = filepath.Abs(file)
		if err != nil {
			logtool.Fetal(err.Error())
		}
	}
	return file
}
