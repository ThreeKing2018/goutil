package golog

import (
	"github.com/ThreeKing2018/goutil/golog/conf"
	"testing"
	"time"
)

func Test_logge(t *testing.T) {
	SetLogger(ZAPLOG,
		conf.WithLogType(conf.LogNormalType),
		conf.WithProjectName("go_love_coin"),
		conf.WithLogType(conf.LogJsontype))

	SetLogLevel(conf.ErrorLevel)
	Debug("this is zap")
	Debug("this is zap")
	SetLogLevel(conf.DebugLevel)
	Debug("this is zap")
	Debug("this is zap")

	time.Sleep(time.Second * 5)

}
