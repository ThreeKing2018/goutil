package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gogs.163.com/feiyu/goutil/golog/conf"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Log struct {
	logger *zap.Logger
	atom   zap.AtomicLevel
}

func parseLevel(level conf.Level) zapcore.Level {
	switch level {
	case conf.DebugLevel:
		return zapcore.DebugLevel
	case conf.InfoLevel:
		return zapcore.InfoLevel
	case conf.WarnLevel:
		return zapcore.WarnLevel
	case conf.ErrorLevel:
		return zapcore.ErrorLevel
	case conf.PanicLevel:
		return zapcore.PanicLevel
	case conf.FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.DebugLevel
}

var encoderConfig = zapcore.EncoderConfig{
	// Keys can be anything except the empty string.
	TimeKey:        "T",
	LevelKey:       "L",
	NameKey:        "N",
	CallerKey:      "C",
	MessageKey:     "M",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func New(opts ...conf.Option) *Log {
	o := &conf.Options{
		Filename:    conf.Filename,
		LogLevel:    conf.LogLevel,
		MaxSize:     conf.MaxSize,
		MaxAge:      conf.MaxAge,
		Stacktrace:  conf.Stacktrace,
		IsStdOut:    conf.IsStdOut,
		ProjectName: conf.ProjectName,
		LogType:     conf.LogNormalType,
	}

	for _, opt := range opts {
		opt(o)
	}

	var writers = []zapcore.WriteSyncer{}
	osfileout := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     o.MaxAge, // days
		LocalTime:  true,
	})
	if o.IsStdOut {
		writers = append(writers, os.Stdout)
	}

	writers = append(writers, osfileout)
	w := zapcore.NewMultiWriteSyncer(writers...)

	atom := zap.NewAtomicLevel()
	atom.SetLevel(parseLevel(o.LogLevel)) //改变日志级别

	var enc zapcore.Encoder
	if o.LogType == conf.LogNormalType {
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		enc = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		//这里控制json 或者不是json 类型
		enc,
		w,
		atom,
	)

	logger := zap.New(
		core,
		zap.AddStacktrace(parseLevel(o.Stacktrace)),
		zap.AddCaller())

	logger = logger.With(zap.String(conf.ProjectKey, o.ProjectName))
	return &Log{logger: logger, atom: atom}

}

func (l *Log) Sync() {
	l.Sync()
}

func (l *Log) SetLogLevel(level conf.Level) {
	l.atom.SetLevel(parseLevel(level))
}

func (l *Log) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(msg)
}
func (l *Log) Info(msg string, fields ...interface{}) {
	l.logger.Info(msg)
}
func (l *Log) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(msg)
}
func (l *Log) Error(msg string, fields ...interface{}) {
	l.logger.Error(msg)
}
func (l *Log) Panic(msg string, fields ...interface{}) {
	l.logger.Panic(msg)
}
func (l *Log) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatal(msg)
}
