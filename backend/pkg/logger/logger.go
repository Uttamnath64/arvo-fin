package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	Development = "DEVELOPMENT"
	Staging     = "STAGING"
	Production  = "PROD"
)

var LogLevel = map[string]zapcore.Level{
	Development: zap.DebugLevel,
	Staging:     zap.InfoLevel,
	Production:  zap.WarnLevel,
}

// logger interface
type LoggerIf interface {
	Debug(topic string, data interface{})
	Info(topic string, data interface{})
	Warn(topic string, data interface{})
	Error(topic string, data interface{})
}

// logger
type Logger struct {
	log *zap.Logger
}

// logger object
func New(logLevel string) *Logger {

	// log level
	level, ok := LogLevel[logLevel]
	if !ok {
		level = zap.DebugLevel
	}

	// // setup zap config
	// config := zap.Config{
	// 	Level:            zap.NewAtomicLevelAt(level),
	// 	Encoding:         "json",
	// 	OutputPaths:      []string{"../logs/app.log"},
	// 	ErrorOutputPaths: []string{"../logs/error.log"},
	// }
	// logger, _ := config.Build()

	// return &Logger{
	// 	log: logger,
	// }

	// configure encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "topic",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// setup console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// setup file encoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// create multi-core with both console and file outputs
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), level),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}), level),
	)

	// create the logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return &Logger{log: logger}

}

// logger Debug log
func (l *Logger) Debug(topic string, arg ...interface{}) {
	l.log.Debug(topic, zap.Any("data", arg))
}

// logger Info log
func (l *Logger) Info(topic string, arg ...interface{}) {
	l.log.Info(topic, zap.Any("data", arg))
}

// logger Warn log
func (l *Logger) Warn(topic string, arg ...interface{}) {
	l.log.Warn(topic, zap.Any("data", arg))
}

// logger Error log
func (l *Logger) Error(topic string, arg ...interface{}) {
	l.log.Error(topic, zap.Any("data", arg))
}

// logger Fatal log
func (l *Logger) Fatal(topic string, arg ...interface{}) {
	l.log.Fatal(topic, zap.Any("data", arg))
}
