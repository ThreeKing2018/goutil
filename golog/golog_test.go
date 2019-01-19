package golog

import (
	"testing"
	"time"

	"github.com/ThreeKing2018/goutil/golog/conf"
)

func Test_logge(t *testing.T) {
	SetLogger(ZAPLOG,
		conf.WithLogType(conf.LogNormalType),
		conf.WithProjectName("go_xxx"),
		conf.WithLogType(conf.LogJsontype))

	SetLogLevel(conf.ErrorLevel)
	Debug("this is zap")
	Debug("this is zap")
	SetLogLevel(conf.DebugLevel)
	Debug("this is zap")
	Debug("this is zap")

	time.Sleep(time.Second * 5)

}
