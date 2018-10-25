package datetime

import (
	"fmt"
	"testing"
	"time"
)

//自定义格式
func TestDateFormat(t *testing.T) {
	fmt.Println(DateFormat(time.Now(), "YYYY/MM/DD"))
}
//获取当前时间
func TestGetNow(t *testing.T) {
	fmt.Println(GetNow())
}
//获取年月日
func TestGetNowDay(t *testing.T) {
	fmt.Println(GetNowDay())
}
//获取时分秒
func TestGetNowHour(t *testing.T) {
	fmt.Println(GetNowHour())
}
//获取昨天的时间
func TestGetBeforeTimesStamp(t *testing.T) {
	fmt.Println(GetBeforeTimesStamp(24*3600*1))
}
//获取明天的时间
func TestGetNextTimesStamp(t *testing.T) {
	fmt.Println(GetNextTimesStamp(24*3600*1))
}
//时间戳转可读时间格式
func TestGetDate(t *testing.T) {
	var tt int64
	tt = 1540437494
	fmt.Println(GetDate(tt))
}