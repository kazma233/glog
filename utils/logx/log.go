package logx

import (
	"glog/config"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// BaseLog Logx 用来记录日志
	BaseLog *zap.Logger
	// ExtLog Logx 用来记录日志
	ExtLog *zap.SugaredLogger
)

const (
	LogMaxSizeMB    = 512
	LogMaxBackupNum = 3
	LogMaxAgeDays   = 30
	DEV             = "DEV"
	PROD            = "PROD"
)

func init() {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl != zapcore.ErrorLevel
	})

	var core zapcore.Core
	if config.Conf().Env == PROD {
		fileErrorDebugging := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "/logs/error.log",
			MaxSize:    LogMaxSizeMB,
			MaxBackups: LogMaxBackupNum,
			MaxAge:     LogMaxAgeDays,
			Compress:   true,
		})
		fileInfoDebugging := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "/logs/info.log",
			MaxSize:    LogMaxSizeMB,
			MaxBackups: LogMaxBackupNum,
			MaxAge:     LogMaxAgeDays,
			Compress:   true,
		})
		fileEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, fileErrorDebugging, highPriority),
			zapcore.NewCore(fileEncoder, fileInfoDebugging, lowPriority),
		)
	} else {
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
	}

	logger := zap.New(core)
	BaseLog = logger
	ExtLog = BaseLog.Sugar()
}

// Debug Debug
func Debug(template string, args ...interface{}) {
	ExtLog.Debugf(template, args...)
}

// Info Info
func Info(template string, args ...interface{}) {
	ExtLog.Infof(template, args...)
}

// Warn Warn
func Warn(template string, args ...interface{}) {
	ExtLog.Warnf(template, args...)
}

// Error Error
func Error(template string, args ...interface{}) {
	ExtLog.Errorf(template, args...)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
