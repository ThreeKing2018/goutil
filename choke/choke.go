package choke

import (
	"os"
	"os/signal"
	"syscall"
	"log"
)

func choke() {

	// 监听退出信号
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	<-c
	log.Println("程序退出")
}
