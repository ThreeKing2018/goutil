package etcdv3

import (
	"context"
	"fmt"
	"github.com/ThreeKing2018/goutil/config/backend/resp"
	goetcd "github.com/coreos/etcd/clientv3"
	"strings"
	"time"
)

type client struct {
	prefix  string
	keysAPI *goetcd.Client
	//waitIndex uint64
	delim string
	// 上次修订版本号
	revision int64
}

func NewClient(endpoint []string, Prefix, Delim string) (*client, error) {
	cfg := goetcd.Config{
		Endpoints:   endpoint,
		DialTimeout: time.Second * 1,
	}

	c, err := goetcd.New(cfg)

	if err != nil {
		return nil, err
	}

	return &client{keysAPI: c, prefix: Prefix, delim: Delim}, nil

}

func (c *client) List(respChan chan *resp.Response) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := c.keysAPI.Get(ctx, c.prefix, goetcd.WithPrefix())
	if err != nil {
		return err
	}

	for _, kv := range r.Kvs {
		respChan <- &resp.Response{
			Key:   convertKey(string(kv.Key), c.prefix, c.delim),
			Value: string(kv.Value),
		}
	}

	close(respChan)
	return nil
}

func (c *client) Watch() <-chan *resp.Response {
	respChan := make(chan *resp.Response, 10)

	go func() {
		rch := c.keysAPI.Watch(
			context.Background(),
			c.prefix,
			goetcd.WithPrefix(),
			goetcd.WithCreatedNotify(),
		)

		for watchResp := range rch {
			for _, ev := range watchResp.Events {
				fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}

	}()

	return respChan
}

//func (c *client) Watch(stop chan struct{}) <-chan *resp.Response {
func (c *client) Watch1() {
	//respChan := make(chan *resp.Response, 10) //加个缓冲区
	//
	//go func() {
	rch := c.keysAPI.Watch(context.Background(), c.prefix)

	for watchResp := range rch {
		fmt.Println("aaaaaaa")
		for _, ev := range watchResp.Events {
			fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	//}()
	//return respChan
}

func convertKey(key, prefix, delim string) string {
	a := strings.TrimPrefix(key, prefix+"/")
	return strings.Replace(a, "/", delim, -1)
}
