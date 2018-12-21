package config

import (
	"fmt"

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
			case a := <-remotechan: //TODO 还有DELETE类型没有判断
				if a.Error != nil {
					continue
				}

				if v.backendName == "file" {
					v.ReadConfig()
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
					fmt.Println(err)
				}

			case <-v.stop:
				return
			}
		}
	}()
}
