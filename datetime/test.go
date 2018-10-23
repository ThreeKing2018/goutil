package datetime

import (
	"fmt"
	"time"
)

func TestDateTime() {
	fmt.Println(GetNow())
	fmt.Println(GetNowDay())
	fmt.Println(GetNowHour())
	fmt.Println(GetBeforeTimesStamp(24*3600*1))
	fmt.Println(GetNextTimesStamp(24*3600*1))
	fmt.Println(DateFormat(time.Now(), "YYYY/MM/DD"))
}