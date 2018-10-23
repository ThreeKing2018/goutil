package main

import (
	"code.sgfoot.com/goutil/convertor"
	"code.sgfoot.com/goutil/datetime"
	"fmt"
)

func main() {
	fmt.Println("测试转换函数:")
	convertor.TestInt64ToBytes()
	convertor.TestBytesToInt64()
	fmt.Println("测试时间函数:")
	datetime.TestDateTime()
}
