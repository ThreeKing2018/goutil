package config

import (
	"github.com/ThreeKing2018/goutil/config/ini"
	"github.com/ThreeKing2018/goutil/config/json"
	"io"
	"strings"
)

type Operater interface {
	SetConfigFile(cnofigfile string)                         //设置配置文件
	ReadConfig(in io.Reader, c map[string]interface{}) error //读取配置文件
	WriteConfig(v interface{}) error                         //写入配置文件
	//WatchConfig() error
	//WatchRemoteConfig()
}

//查找配置文件类型 并初始化 ini yaml yml
func getType(cfgtype string) (Operater, error) {
	//转换成小写
	cfgtype = strings.ToLower(cfgtype)

	if !stringInSlice(cfgtype, supportedExts) {
		return nil, UnsupportedConfigError(cfgtype)
	}

	var a Operater

	switch cfgtype {
	case "json":
		a = json.NewConf()
	case "ini":
		a = ini.NewConf()

	}

	return a, nil
}
