package convertor

import (
	"encoding/binary"
	"reflect"
	"fmt"
	"strconv"
)


//byte 转 int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
//string to int
func StringToInt(value string) (i int) {
	i, _ = strconv.Atoi(value)
	return
}
//int to int64
func IntToInt64(value int) int64 {
	i, _ := strconv.ParseInt(string(value), 10, 64)
	return i
}
// convert any numeric value to int64
// 任意类型转int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

