package gotime

import (
	"fmt"
	"strings"
	"time"
)

type GoTime struct {
	Location time.Location
}

//实例
func New() *GoTime {
	return &GoTime{}
}

//获取当前时间 年-月-日 时:分:秒
func (gt *GoTime) Now() string {
	return gt.NowTime().Format(TT)
}

//获取当前时间戳
func (gt *GoTime) NowUnix() int64 {
	return gt.NowTime().Unix()
}

//获取当前时间Time
func (gt *GoTime) NowTime() time.Time {
	return time.Now().In(Location)
}

//获取年月日
func (gt *GoTime) GetYmd() string {
	return gt.NowTime().Format(YMD)
}

//获取时分秒
func (gt *GoTime) GetHms() string {
	return gt.NowTime().Format(HMS)
}

//获取当天的开始时间, eg: 2018-01-01 00:00:00
func (gt *GoTime) NowStart() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, Location)
	return tm.Format(TT)
}

//获取当天的结束时间, eg: 2018-01-01 23:59:59
func (gt *GoTime) NowEnd() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, Location)
	return tm.Format(TT)
}

//当前时间 减去 多少秒
func (gt *GoTime) Before(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()-beforeSecond, 0).Format(TT)
}

//当前时间 加上 多少秒
func (gt *GoTime) Next(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()+beforeSecond, 0).Format(TT)
}

//2006-01-02T15:04:05Z07:00 转 时间戳
func (gt *GoTime) RfcToUnix(layout string) int64 { //转化所需模板
	tm, err := time.ParseInLocation(time.RFC3339, layout, Location) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return int64(0)
	}
	return tm.Unix()
}

//2006-01-02 15:04:05 转 时间戳
func (gt *GoTime) ToUnix(layout string) int64 {
	theTime, _ := time.ParseInLocation(TT, layout, Location)
	return theTime.Unix()
}

//获取RFC3339格式
func (gt *GoTime) GetRFC3339() string {
	return gt.NowTime().Format(time.RFC3339)
}

//转换成RFC3339格式
func (gt *GoTime) ToRFC3339(layout string) string {
	tm, err := time.ParseInLocation(TT, layout, Location)
	if err != nil {
		return ""
	}
	return tm.Format(time.RFC3339)
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
func (gt *GoTime) Format(t time.Time, format string) string {
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
