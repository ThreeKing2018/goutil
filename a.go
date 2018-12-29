package main

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/config/backend/etcdv3"
	"github.com/ThreeKing2018/goutil/config/backend/resp"
)

func main() {
	c, err := etcdv3.NewClient([]string{"127.0.0.1:2379"}, "root", ".")
	if err != nil {
		fmt.Println(err)
	}

	respChan := make(chan *resp.Response, 10)

	err = c.List(respChan)
	fmt.Println("asskdsdsdsddsjndcd")
	fmt.Println(err)

	fmt.Println("asskdsdsdsddsjndcd")
	fmt.Println(<-respChan)

	fmt.Println("asskdsdsdsd")
	for {
		select {
		case v := <-respChan:
			fmt.Println("bbbbbbb")
			fmt.Println("a")
			fmt.Println(v)

		}
	}
}
