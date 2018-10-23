package config

import (
	"bytes"
	"fmt"
	"gogs.163.com/feiyu/goutil/files"
	"log"
	"os"
	"path"
	//"gogs.163.com/feiyu/goutil/config/backends"
	"github.com/coreos/etcd/client"
	"gogs.163.com/feiyu/goutil/config/backend"
	"gogs.163.com/feiyu/goutil/config/backend/resp"
	"sync"
	"time"
)

//接口添加able 组合 viperable
type Viperable interface {
	//配置文件的初始化
	SetDefault(key string, value interface{})
	SetConfig(cfgfile, cfgtype string, cfgpath ...string)
	SetKeyDelim(delim string)

	//WatchConfig(remoteCfg *backend.Config) error
	WatchConfig() error
	Stop()
	//Getdefault() map[string]interface{}
	//Getconfig() map[string]interface{}

	Getconfig() map[string]interface{}
	//Get方法
	Getvalue

	//AddConfigPath(cfgpath string)
	//Operater

	//对配置文件的操作
	ReadConfig() error
	WriteConfig() error

	ReadRemoteConfig() error

	SetFunc(fn func(key string, value interface{}) error)

	//设置本地配置文件
	LocalConfig() *backend.Config

	//设置远程接口
	EtcdConfig(prefix string, endpoint []string)

	//远程配置
	//AddRemoteProvider(provider, endpoint, path string) error

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
	config map[string]interface{}
	//override       map[string]interface{}
	defaults map[string]interface{}
	//kvstore        map[string]interface{}
	//pflags         map[string]FlagValue
	stop chan struct{}
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

func (v *viper) LocalConfig() *backend.Config {

	if v.configFile == "" {
		panic(NotFoundConfigError(v.configName))
	}

	a := &backend.Config{
		Delim:   v.keyDelim,
		Backend: "file",
	}
	v.backendInit(a)

	return a
}

func (v *viper) EtcdConfig(prefix string, endpoint []string) {
	a := &backend.Config{
		Delim:    v.keyDelim,
		Backend:  "etcd",
		Prefix:   prefix,
		Endpoint: endpoint,
	}

	v.backendInit(a)

}

func (v *viper) backendInit(a *backend.Config) {
	v.backendName = a.Backend

	a.ConfigFiles = v.configFile
	var err error
	v.backend, err = backend.New(a)
	if err != nil {
		log.Println(err)
	}
}

func (v *viper) getBackendName() string {
	return v.backendName
}

func (v *viper) SetFunc(fn func(key string, value interface{}) error) {
	v.fn = fn
}

func (v *viper) Stop() {
	v.stop <- struct{}{}
}

// SetConfigName 设置配置文件的名称
func (v *viper) SetConfig(cfgfile, cfgtype string, cfgpath ...string) {
	log.Println(fmt.Sprintf("配置文件:%s", cfgfile))

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

}

func (v *viper) SetKeyDelim(delim string) {
	log.Println(fmt.Sprintf("设置key的分隔符 %s", delim))
	if delim != "" {
		v.keyDelim = delim
	}
}

// SetDefault 注册默认值
func (v *viper) SetDefault(key string, value interface{}) {
	v.defaults[key] = value
	//log.Println(fmt.Sprintf("key:%s,value:%s",key,value))
}

//读取配置失败 具体的操作可以交给 调用者
func (v *viper) ReadConfig() error {
	file, err := files.ReadFile(v.configFile)
	if err != nil {
		return configReadError(v.configFile)
	}

	err = v.operating.ReadConfig(bytes.NewReader(file), v.config)
	if err != nil {
		return configParseError{err}
	}

	return nil
}

func (v *viper) ReadRemoteConfig() error {
	//读取远程配置
	var err error
	respChan := make(chan *resp.Response, 10) //添加缓存 让etcd尽快处理完

	go func() {
		for {
			select {
			case a, ok := <-respChan:
				if !ok {
					//读取完成关闭通道，写入配置文件到本地
					err = v.WriteConfig()
					if err != nil {
						fmt.Println(err, "关闭通道")
					}
					return
				}
				//fmt.Println(a)
				err = v.Set(a.Key, a.Value)

			}

		}
	}()

	err = v.backend.List(respChan)
	if err != nil {
		return err
	}

	return nil
}

func (v *viper) WriteConfig() error {
	return v.operating.WriteConfig(v.config)
}

//这里的stop 以后可以换成context,,  string, map，[]string  []int 都能自动转换,其他类型需要自定义fn函数
func (v *viper) WatchConfig() error {

	remotechan := v.backend.Watch(v.stop)
	//var a backends.Response
	go func() {
		for {
			select {
			case a := <-remotechan: //TODO 还有DELETE类型没有判断
				if a.Error != nil {
					fmt.Println("err1", a.Error)
					continue
				}

				if v.getBackendName() == "file" {
					v.ReadConfig()
					continue
				}

				//这里是其他的backend

				v.Set(a.Key, a.Value)

				//保存到本地
				err := v.WriteConfig()
				if err != nil {
					fmt.Println(err)
				}

			case <-v.stop:
				return
			}
		}
	}()
	//r.WatchConfig(v)

	return nil
}

type Client struct {
	client   client.KeysAPI
	prefix   string
	stopChan chan struct{}
	close    bool
}

// NewEtcdClient returns an *etcd.Client with a connection to named machines.
func NewClient(endpoint []string, Prefix string) (*Client, error) {
	cfg := client.Config{
		Endpoints:               endpoint,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second * 1,
	}

	var kapi client.KeysAPI

	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	kapi = client.NewKeysAPI(c)
	return &Client{client: kapi, prefix: Prefix}, nil
}

func (v *viper) Getdefault() map[string]interface{} {
	return v.defaults
}

func (v *viper) Getconfig() map[string]interface{} {
	return v.config
}
