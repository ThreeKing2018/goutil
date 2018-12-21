package gotime

import (
	"fmt"
	"testing"
)

func TestGoTime_Now(t *testing.T) {
	tm := New().Now()
	fmt.Println(tm)
}
func TestGoTime_Before(t *testing.T) {
	tm := New().Before(3600)
	fmt.Println(tm)
}
func TestGoTime_Next(t *testing.T) {
	tm := New().Next(3600)
	fmt.Println(tm)
}
func TestGoTime_GetYmd(t *testing.T) {
	tm := New().GetYmd()
	fmt.Println(tm)
}
func TestGoTime_NowTime(t *testing.T) {
	tm := New().NowTime()
	fmt.Println(tm)
}
func TestGoTime_RfcToUnix(t *testing.T) {
	tm := New().RfcToUnix("2018-12-12T15:28:16+08:00")
	fmt.Println(tm)
}
func TestGoTime_GetRFC3339(t *testing.T) {
	tm := New().GetRFC3339()
	fmt.Println(tm)
}
func TestGoTime_ToRFC3339(t *testing.T) {
	tm := New().ToRFC3339(New().Now())
	fmt.Println(tm)
}
