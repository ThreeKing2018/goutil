package etcdv3

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/config/backend/resp"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	c, err := NewClient([]string{"127.0.0.1:2379"}, "root", ".")
	if err != nil {
		t.Log(err)
	}

	respChan := make(chan *resp.Response, 10)

	go func() {
		for {
			select {
			case v, ok := <-respChan:
				if ok {
					fmt.Println(v)
				} else {
					fmt.Println("关闭chan")
					return
				}

			}
		}
	}()

	err = c.List(respChan)
	fmt.Println(err)

	time.Sleep(1 * time.Second)

	//c.Watch()

}
