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
		conf.WithLogType(conf.LogJsontype),
		conf.WithFilename("log.txt"),
		conf.WithIsStdOut(false))

	SetLogLevel(conf.ErrorLevel)
	Debug("this is zap")
	Debug("this is zap")
	SetLogLevel(conf.DebugLevel)
	Debug("this is zap")
	Debug("this is zap")
	Infow("aa", "aaa",100)

	time.Sleep(time.Second * 5)

}
