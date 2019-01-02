package file

import (
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

	//inode
	watcher, err := fsnotify.NewWatcher()
	//监视配置文件inode 出错了,退出程序
	if err != nil {
		panic(err)
	}
	go func() {
		respdata := &resp.Response{
			Error: nil,
		}

		for {
			select {
			case event,ok := <-watcher.Events:
				if !ok {
					return
				}

				if  event.Op&fsnotify.Create == fsnotify.Create ||
					event.Op&fsnotify.Rename == fsnotify.Rename ||
					event.Op&fsnotify.Write == fsnotify.Write{

					//需要读取配置文件
					//通过chan通知
					err = watcher.Add(c.configFile)
					respChan <- respdata
				}
			case err,ok := <-watcher.Errors:
				if !ok {
					return
				}
				respdata.Error = err
				respChan <- respdata
			case <-stop:
				watcher.Close()
				close(respChan)
			}

		}


	}()

	watcher.Add(c.configFile)

	return respChan
}
