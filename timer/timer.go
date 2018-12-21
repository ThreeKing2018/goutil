package timer

import (
	"fmt"
	"log"
	"time"

	"github.com/yezihack/gotime"
)

func Start() {
	//开个协程
	go StartTicker()
}
func StartTicker() {
	log.Println("定时器启动")
	//new一个定时器(设置一个间隔时间)
	ticker := time.NewTicker(time.Duration(1) * time.Second * 3)
	//使用一个死循环
	for range ticker.C {
		//todo
		fmt.Println("定时器正在运行..." + gotime.NewGoTime().Now())
	}
}
