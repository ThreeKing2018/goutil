package golog

import "github.com/ThreeKing2018/goutil/golog/conf"

//使用string是为了减少使用Spintf
type ILog interface {
	//new
	//普通日志,如果有args，需要格式化
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Panic(string, ...interface{})
	Fatal(string, ...interface{})
	SetLogLevel(level conf.Level)
	Sync()
	//需要格式化日志 ，最后一个是context
	//Debugf(string, ...interface{})
	//Infof(string, ...interface{})
	//Warnf(string, ...interface{})
	//Errorf(string, ...interface{})
	//Panicf(string, ...interface{})
	//Fatalf(string, ...interface{})
}
