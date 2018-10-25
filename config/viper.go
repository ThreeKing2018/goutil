package config

import (
	"github.com/ThreeKing2018/goutil/config/backend"
	"os"
	"path"
	"sync"
)

type Viperable interface {
	SetDefault(key string, value interface{}) Viperable
	SetKeyDelim(delim string) Viperable
	SetConfig(cfgfile, cfgtype string, cfgpath ...string) Viperable
	SetFunc(fn func(key string, value interface{}) error) Viperable
	RemoteConfig(prefix string, endpoint []string) Viperable

	//配置文件的操作
	writeConfig() error
	ReadConfig() error

	//配置文件动态加载
	WatchConfig() error
	Stop()
}

type viper struct {
	// Delimiter that separates a list of keys
	// used to access a nested value in one go
	keyDelim string

	// A set of paths to look for the config file in
	configPaths []string

	configFile string //配置文件的绝对路径
	configName string //配置文件的名称
	configType string

	//remoteProviders []*backends.ProviderConfig
	operating Operater
	backend   backend.StoreClient

	// 配置值相关的
	config   map[string]interface{}
	defaults map[string]interface{}
	stop     chan struct{}
	//TODO改怎么解释这个函数的作用大呢
	fn func(key string, value interface{}) error

	backendName string
	mu          sync.RWMutex
}

func New() Viperable {
	v := new(viper)

	//key的分隔符，分割成列表
	v.keyDelim = "."
	v.config = make(map[string]interface{})
	v.defaults = make(map[string]interface{})
	v.stop = make(chan struct{}, 0)
	v.fn = func(key string, value interface{}) error {
		return nil
	}

	return v
}

// 注册默认值
func (v *viper) SetDefault(key string, value interface{}) Viperable {
	v.defaults[key] = value
	return v
}

func (v *viper) SetKeyDelim(delim string) Viperable {
	if delim != "" {
		v.keyDelim = delim
	}

	return v
}

// SetConfigName 设置配置文件的名称
func (v *viper) SetConfig(cfgfile, cfgtype string, cfgpath ...string) Viperable {
	if cfgtype == "" {
		panic("配置类型不能为空")
	}

	if cfgfile == "" {
		panic("配置文件不能为空")
	}

	v.configName = cfgfile //只是保存用户传递过来名称

	var err error
	// 循环path 判断配置文件 存在(IsExist true) 赋值 configfile = _cfgfile,
	// 找不到则configfile:=cfgfile
	cfgpath = append(cfgpath, ".")
	for _, i := range cfgpath {
		_cfgfile := path.Join(i, cfgfile)
		_, err := os.Stat(_cfgfile)
		if err == nil { //TODO 使用os.IsExist(err) 每次都是都返回 false
			v.configFile = _cfgfile
			continue

		}
	}

	if v.configFile == "" {
		v.configFile = v.configName
	}

	//根据类型初始化配置
	v.operating, err = getType(cfgtype)
	if err != nil {
		panic(err)
	}

	v.configType = cfgtype

	//赋值给Operater接口
	v.operating.SetConfigFile(v.configFile)

	return v

}

func (v *viper) SetFunc(fn func(key string, value interface{}) error) Viperable {
	v.fn = fn
	return v
}

func (v *viper) writeConfig() error {
	return v.operating.WriteConfig(v.config)
}
