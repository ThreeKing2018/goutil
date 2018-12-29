package file

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/config/backend/resp"
	"github.com/fsnotify/fsnotify"
)

type client struct {
	configFile string
}

func NewClient(configFile string) (*client, error) {
	return &client{
		configFile: configFile,
	}, nil
}

func (c *client) List(respChan chan *resp.Response) error {
	return nil
}

func (c *client) Watch(stop chan struct{}) <-chan *resp.Response {
	respChan := make(chan *resp.Response, 10) //加个缓冲区

	go func() {
		//inode
		watcher, err := fsnotify.NewWatcher()
		//监视配置文件inode 出错了,退出程序
		if err != nil {
			panic(err)
		}

		watcher.Add(c.configFile)

		go func() {
			<-stop
			watcher.Close()
		}()

		respdata := &resp.Response{
			Error: nil,
		}

		for {
			fmt.Println("a")
			select {
			case event := <-watcher.Events:
				fmt.Println("event", event)
				if event.Op&fsnotify.Remove == fsnotify.Remove ||
					event.Op&fsnotify.Rename == fsnotify.Rename ||
					event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Create == fsnotify.Create {
					watcher.Remove(c.configFile)
					watcher.Close()
					watcher.Add(c.configFile)
					fmt.Println("aaaaaaaaaaaaaaaa")
					//需要读取配置文件
					//通过chan通知
					respChan <- respdata
					break
				}

			case err := <-watcher.Errors:
				respdata.Error = err
				respChan <- respdata
			}

		}
	}()

	return respChan
}
