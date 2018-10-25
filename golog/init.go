package golog

import (
	"github.com/ThreeKing2018/goutil/golog/conf"
	"github.com/ThreeKing2018/goutil/golog/plugins/zaplog"
)

//默认
var l ILog = zaplog.New()

type backend uint8

const (
	ZAPLOG backend = iota
	DlOG
)

//设置
func SetLogger(b backend, opts ...conf.Option) {
	switch b {
	case ZAPLOG:
		l = zaplog.New(opts...)
	case DlOG:

	}
}

//目前只有zap生效
func SetLogLevel(level conf.Level) {
	l.SetLogLevel(level)
}

//目前只有zap生效
func Sync() {
	l.Sync()
}

//普通日志
func Debug(msg string, args ...interface{}) {
	l.Debug(msg, args...)
}
func Info(msg string, args ...interface{}) {
	l.Info(msg, args...)
}
func Warn(msg string, args ...interface{}) {
	l.Warn(msg, args...)
}
func Error(msg string, args ...interface{}) {
	l.Error(msg, args...)
}
func Panic(msg string, args ...interface{}) {
	l.Panic(msg, args...)
}
func Fatal(msg string, args ...interface{}) {
	l.Fatal(msg, args...)
}
