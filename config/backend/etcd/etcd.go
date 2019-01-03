package etcd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ThreeKing2018/goutil/config/backend/resp"
	goetcd "go.etcd.io/etcd/client"
	)

type client struct {
	prefix    string
	keysAPI   goetcd.KeysAPI
	waitIndex uint64
	delim     string
}

func NewClient(endpoint []string, Prefix, Delim string) (*client, error) {
	cfg := goetcd.Config{
		Endpoints:               endpoint,
		HeaderTimeoutPerRequest: time.Second * 1,
	}

	c, err := goetcd.New(cfg)

	if err != nil {
		return nil, err
	}
	keysAPI := goetcd.NewKeysAPI(c)

	return &client{keysAPI: keysAPI, prefix: Prefix, delim: Delim}, nil

}

/*	Get(key string) ([]byte, error)

	// List retrieves all keys and values under a provided key.
	List(key string) (KVPairs, error)

	// Set sets the provided key to value.
	Set(key string, value []byte) error

	// Watch monitors a K/V store for changes to key.
	Watch(key string, stop chan bool) <-chan *Response*/
//func (c *client) Get(key string) ([]byte,error) {
//	return nil,nil
//}
//
//
//func (c *client) Set(key, value []byte) error {
//	return nil
//}

// nodeWalk recursively descends nodes, updating vars.
func (c *client) nodeWalk(node *goetcd.Node, respChan chan *resp.Response) error {
	if node != nil {
		key := convertKey(node.Key, c.prefix, c.delim)
		if !node.Dir {
			respChan <- &resp.Response{Key: key, Value: node.Value}
			//vars[key] = node.Value
		} else {
			for _, node := range node.Nodes {
				c.nodeWalk(node, respChan)
			}
		}
	}
	return nil
}

func (c *client) List(respChan chan *resp.Response) error {
	resp, err := c.keysAPI.Get(context.Background(), c.prefix, &goetcd.GetOptions{
		Recursive: true,
		Sort:      true,
		Quorum:    true,
	})
	if err != nil {
		return err
	}
	err = c.nodeWalk(resp.Node, respChan)
	if err != nil {
		return err
	}

	close(respChan)
	return nil
}

func (c *client) Watch(stop chan struct{}) <-chan *resp.Response {
	respChan := make(chan *resp.Response, 10) //加个缓冲区
	go func() {
		watcher := c.keysAPI.Watcher(c.prefix, &goetcd.WatcherOptions{
			Recursive: true,
		})
		ctx, cancel := context.WithCancel(context.TODO())

		go func() {
			<-stop
			cancel()
		}()

		respdata := &resp.Response{
			Error: nil,
		}

		for {
			var resp *goetcd.Response
			var err error
			resp, err = watcher.Next(ctx)
			if err != nil {
				respdata.Error = err
				respChan <- respdata
				time.Sleep(time.Second * 1)
				continue
			}

			respdata.Action = resp.Action

			switch resp.Action {
			case "set", "update":
				respdata.Key = convertKey(resp.Node.Key, c.prefix, c.delim)
				respdata.Value = resp.Node.Value

			case "delete":
				respdata.Key = convertKey(resp.Node.Key, c.prefix, c.delim)

			default:
				respdata.Error = fmt.Errorf("没有发现的action")

			}

			respChan <- respdata

		}

	}()
	return respChan
}

func convertKey(key, prefix, delim string) string {
	a := strings.TrimPrefix(key, prefix+"/")
	return strings.Replace(a, "/", delim, -1)
}
