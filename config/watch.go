package config

import "fmt"

func (v *viper) Stop() {
	v.stop <- struct{}{}
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

				if v.backendName == "file" {
					v.ReadConfig()
					continue
				}

				//这里是其他的backend

				v.Set(a.Key, a.Value)

				//保存到本地
				err := v.writeConfig()
				//TODO 这里的错误怎么处理呢
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
