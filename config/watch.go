package config

import (
	"time"

	"github.com/ThreeKing2018/goutil/golog"
)

func (v *viper) Stop() {
	v.stop <- struct{}{}
}

//这里的stop 以后可以换成context,,  string, map，[]string  []int 都能自动转换,其他类型需要自定义fn函数
func (v *viper) WatchConfig() {
	remotechan := v.backend.Watch(v.stop)
	//var a backends.Response
	go func() {
		for {
			select {
			case a,ok := <-remotechan: //TODO 还有DELETE类型没有判断
			if !ok {
				golog.Error("watch配置出问题了,重启开启watch")
				time.Sleep(1*time.Second)
				v.Stop()
				go v.WatchConfig()
				return
			}
				if a.Error != nil {
					golog.Error(a.Error)
					continue
				}

				if v.backendName == "file" {
					v.ReadConfig()
					//v.Stop()
					continue
				}

				//这里是其他的backend
				err := v.Set(a.Key, a.Value)
				if err != nil {
					golog.Error(err.Error())
				}

				//保存到本地
				err = v.writeConfig()
				if err != nil {
					golog.Error(err.Error())
				}

			case <-v.stop:
				return
			}
		}
	}()
}
