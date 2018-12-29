package config

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cast"
)

/*

func Test_JsonConf(t *testing.T) {
	a := GetInstance()
	a.ParseConf("../cmd/server.json","json")
	fmt.Println(a.Cfg.Version)
	a.Cfg.Version = "1.121"
	a.Cfg.Mysql.DBpasswd = "aa"
	fmt.Println(a.Cfg.Version)

	a.C.SaveFile()
	fmt.Println(a.Cfg.Version)
	time.Sleep(10*time.Second)
	a.C.Reload()
	fmt.Println(a.Cfg.Version)
	t.Log(a.Cfg)
}

*/

func Test_viper(t *testing.T) {
	var err error
	v := New()
	//v.SetKeyDelim("-")
	v.SetConfig("viper.json", "json", "/etc", "/home", ".")
	v.SetDefault("debug", false)
	v.SetDefault("port1", 7777)

	//开启动态配置文件
	//remoteCfg := &backend.Config{
	//	Backend:"file",
	//	Endpoint:[]string{"http://10.211.55.4:2378"},
	//	Prefix:"/test/aaa",
	//}
	//remoteCfg := v.LocalConfig()

	fn := func(key string, value interface{}) error {

		//int bool
		switch key {
		case "debug":
			value, err = cast.ToBoolE(value)
			return err
		case "port":
			value, err = cast.ToIntE(value)
			return err
		}
		return nil
	}

	v.SetFunc(fn)
	//remoteCfg := &backend.Config{Backend:"file"}
	err = v.ReadConfig()
	v.WatchConfig()
	//err = v.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//for {
	//	fmt.Println(v.GetBool("debug"))
	//}
	//for {
	//	time.Sleep(3*time.Second)
	//	fmt.Println(v.GetInt("asasa"))
	//	fmt.Println(v.GetBool("debug"))
	//	//fmt.Println(v.GetStringMap("key4")["a"])
	//	//fmt.Println(v.GetIntSlice("key3"))
	//	//a := v.GetIntSlice("key3")
	//	//fmt.Println(reflect.TypeOf(a))
	//	//fmt.Println(len(a))
	//
	//
	//}
	////fmt.Println(v.Getconfig())
	time.Sleep(100 * time.Hour)
	v.Stop()

	//fmt.Println(v.Getdefault())

}
