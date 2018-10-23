package datetime

import (
	"time"
	"strings"
	"fmt"
)

var ChinaLocation = time.FixedZone("Asia/Shanghai", 8*60*60)
var format = "2006-01-02 15:04:05"

//获取当前时间 格式: 2018-11-11 00:00:00
func GetNow() (string) {
	return time.Now().In(ChinaLocation).Format(format)
}
//获取当前年月日时间
func GetNowDay() string {
	return time.Now().In(ChinaLocation).Format("2006-01-02")
}
//获取当前的时分秒时间
func GetNowHour() string {
	return time.Now().In(ChinaLocation).Format("060102")
}
//获取减去的时间戳
func GetBeforeTimesStamp(beforeSecond int) string {
	return time.Unix(time.Now().In(ChinaLocation).Unix() - int64(beforeSecond),0).Format(format)
}
//获取追加的时间戳
func GetNextTimesStamp(nextSecond int) string {
	return time.Unix(time.Now().In(ChinaLocation).Unix() + int64(nextSecond),0).Format(format)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func DateFormat(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}