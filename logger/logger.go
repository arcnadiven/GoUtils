package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	std ILogger = NewLogger("")
)

type ILogger interface {
	Level() zapcore.Level
	Log(lvl zapcore.Level, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Logf(lvl zapcore.Level, template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Logw(lvl zapcore.Level, msg string, keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Logln(lvl zapcore.Level, args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	DPanicln(args ...interface{})
	Panicln(args ...interface{})
	Fatalln(args ...interface{})
	Sync() error
}

func NewLogger(logPath string) ILogger {
	logger := io.Writer(nil)
	logger = os.Stdout
	callerSkip := 1 // stdout 默认为1. 输出文件 默认为0
	if logPath != "" {
		logger = &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 3,
			LocalTime:  true,
			Compress:   true,
		}
		callerSkip = 0
	}

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.ConsoleSeparator = " "
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(logger),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)
	sugar := zap.New(core).
		WithOptions(zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(callerSkip)).
		Sugar()

	return sugar
}

func Debugf(template string, args ...interface{})  { std.Debugf(template, args...) }
func Infof(template string, args ...interface{})   { std.Infof(template, args...) }
func Warnf(template string, args ...interface{})   { std.Warnf(template, args...) }
func Errorf(template string, args ...interface{})  { std.Errorf(template, args...) }
func DPanicf(template string, args ...interface{}) { std.DPanicf(template, args...) }
func Panicf(template string, args ...interface{})  { std.Panicf(template, args...) }
func Fatalf(template string, args ...interface{})  { std.Fatalf(template, args...) }
func Debugln(args ...interface{})                  { std.Debugln(args...) }
func Infoln(args ...interface{})                   { std.Infoln(args...) }
func Warnln(args ...interface{})                   { std.Warnln(args...) }
func Errorln(args ...interface{})                  { std.Errorln(args...) }
func DPanicln(args ...interface{})                 { std.DPanicln(args...) }
func Panicln(args ...interface{})                  { std.Panicln(args...) }
func Fatalln(args ...interface{})                  { std.Fatalln(args...) }
