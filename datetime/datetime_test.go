package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	fmt.Println(DateFormat(time.Now(), "YYYY/MM/DD"))
}
func TestGetNow(t *testing.T) {
	fmt.Println(GetNow())
}
func TestGetNowDay(t *testing.T) {
	fmt.Println(GetNowDay())
}
func TestGetNowHour(t *testing.T) {
	fmt.Println(GetNowHour())
}
func TestGetBeforeTimesStamp(t *testing.T) {
	fmt.Println(GetBeforeTimesStamp(24*3600*1))
}
func TestGetNextTimesStamp(t *testing.T) {
	fmt.Println(GetNextTimesStamp(24*3600*1))
}