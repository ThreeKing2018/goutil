package config

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"sync"

	"github.com/ThreeKing2018/goutil/config/backend"
	"github.com/ThreeKing2018/goutil/config/backend/resp"
)

var ErrExit = errors.New("退出配置模块")

type Viperable interface {
	SetDefault(key string, value interface{}) Viperable
	SetKeyDelim(delim string) Viperable
	SetConfig(cfgfile, cfgtype string, cfgpath ...string) Viperable
	SetFunc(fn func(key string, value interface{}) error) Viperable
	SetRemote(prefix, remoteType string, endpoint []string) Viperable
	//
	////配置文件的操作
	writeConfig() error
	ReadConfig() error
	//
	////配置文件动态加载
	WatchConfig()
	Stop()

	//Get方法
	Getvalue
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
	remoteConf *backend.Config
	config     map[string]interface{}
	defaults   map[string]interface{}
	stop       chan struct{}
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
	v.remoteConf = nil
	v.backendName = "file"

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

func (v *viper) SetRemote(prefix, remoteType string, endpoint []string) Viperable {
	v.remoteConf = &backend.Config{
		Delim:    v.keyDelim,
		Backend:  remoteType,
		Prefix:   prefix,
		Endpoint: endpoint,
	}
	v.backendName = remoteType
	return v
}

func (v *viper) ReadConfig() error {
	if v.remoteConf != nil {
		//获取远程配置
		return v.readRemoteConfig()
	}
	//读取本地配置
	return v.localReadConfig()
}

//读取配置失败 具体的操作可以交给 调用者
func (v *viper) localReadConfig() error {
	file, err := readFile(v.configFile)
	if err != nil {
		return configReadError(v.configFile)
	}
	err = v.operating.ReadConfig(bytes.NewReader(file), v.config)
	if err != nil {
		return configParseError{err}
	}

	v.remoteConf = &backend.Config{
		Backend:     "file",
		ConfigFiles: v.configFile,
	}
	v.backend, err = backend.New(v.remoteConf)
	if err != nil {
		return err
	}
	v.remoteConf = nil
	return nil
}

func (v *viper) readRemoteConfig() error {
	v.remoteConf.ConfigFiles = v.configFile

	var err error
	v.backend, err = backend.New(v.remoteConf)
	if err != nil {
		return err
	}
	//读取远程配置

	respChan := make(chan *resp.Response, 10) //添加缓存 让etcd尽快处理完

	//TODO 这里错误需要错误下
	go func()  {
		for {
			select {
			case a, ok := <-respChan:
				if !ok {
					//读取完成关闭通道，写入配置文件到本地
					err = v.writeConfig()
					if err != nil {
						return
					}
					return
				}
				err = v.Set(a.Key, a.Value)
			case <-v.stop:
				{
					return
				}
			}
		}
	}()

	err = v.backend.List(respChan)
	if err != nil {
		return err
	}

	return nil
}

//打开文件计算文件大小，然后创建n+ytes.MinRead (512)大小的buf,
func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// It's a good but not certain bet that FileInfo will tell us exactly how much to
	// read, so let's try it but be prepared for the answer to be wrong.
	var n int64

	if fi, err := f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	// As initial capacity for readAll, use n + a little extra in case Size is zero,
	// and to avoid another allocation after Read has filled the buffer.  The readAll
	// call will read into its allocated internal buffer cheaply.  If the size was
	// wrong, we'll either waste some space off the end or reallocate as needed, but
	// in the overwhelmingly common case we'll get it just right.
	return readAll(f, n+bytes.MinRead)
}

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.

//使用_, err = buf.ReadFrom(r)读取所有数据,返回buf
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
