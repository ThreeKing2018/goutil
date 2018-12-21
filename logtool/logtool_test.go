package logtool

import "testing"

func Test_Trace(t *testing.T) {
	Trace("我的跟踪:%s", "Test_Trace")
}
func Test_Debug(t *testing.T) {
	Debug("我的跟踪:%s", "Test_Debug")
}
func Test_Info(t *testing.T) {
	Info("我的跟踪:%s", "Test_Info")
}
func Test_Warn(t *testing.T) {
	Warn("我的跟踪:%s", "Test_Warn")
}
func Test_Error(t *testing.T) {
	Error("我的跟踪:%s", "Test_Error")
}
func Test_Fetal(t *testing.T) {
	Fetal("我的跟踪:%s", "Test_Fetal")
}
func Test_SetLevel(t *testing.T) {
	//如果设置高级别的话,低级别的错误将不会输出,后面warn, error不会被打印
	SetLevelWithDefault("aa", "F")
	Trace("我的跟踪:%s", "test")
	Warn("我不会被打印")
	Error("我不会被打印")
}
