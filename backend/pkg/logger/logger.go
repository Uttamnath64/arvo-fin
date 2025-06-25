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

// NewTest creates a logger that logs only to the console (for testing)
func NewTest(logLevel string) *Logger {
	// log level
	level, ok := LogLevel[logLevel]
	if !ok {
		level = zap.DebugLevel
	}

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

	// console encoder and core only
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	// build logger
	logger := zap.New(consoleCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return &Logger{log: logger}
}

func toZapFields(fields []interface{}) []zap.Field {
	var zapFields []zap.Field
	if len(fields)%2 != 0 {
		// odd number of args, fallback
		zapFields = append(zapFields, zap.Any("fields", fields))
		return zapFields
	}

	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			// if key is not string, fallback to printing whole slice as data
			return []zap.Field{zap.Any("fields", fields)}
		}
		zapFields = append(zapFields, zap.Any(key, fields[i+1]))
	}
	return zapFields
}

func (l *Logger) Debug(topic string, fields ...interface{}) {
	zapFields := toZapFields(fields)
	l.log.Debug(topic, zapFields...)
}

func (l *Logger) Info(topic string, fields ...interface{}) {
	zapFields := toZapFields(fields)
	l.log.Info(topic, zapFields...)
}

func (l *Logger) Warn(topic string, fields ...interface{}) {
	zapFields := toZapFields(fields)
	l.log.Warn(topic, zapFields...)
}

func (l *Logger) Error(topic string, fields ...interface{}) {
	zapFields := toZapFields(fields)
	l.log.Error(topic, zapFields...)
}

func (l *Logger) Fatal(topic string, fields ...interface{}) {
	zapFields := toZapFields(fields)
	l.log.Fatal(topic, zapFields...)
}
