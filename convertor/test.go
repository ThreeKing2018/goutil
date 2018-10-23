package convertor

import (
	"fmt"
)

func TestInt64ToBytes() {
	var i int64
	i = 10001
	rs := Int64ToBytes(i)
	fmt.Println(rs)
}
func TestBytesToInt64() {
	var b []byte
	b = []byte{0,0,0, 0, 0, 0, 39, 17}
	_i64 := BytesToInt64(b)
	fmt.Println(_i64)
}