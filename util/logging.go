// Package util util関数
package util

import (
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel    zapcore.Level
	appVersion  string
	appRevision string
	localLogger *zap.SugaredLogger
	once        sync.Once
)

const (
	versionKey  = "version"
	revisionKey = "revision"
)

// Config ログ設定
type Config struct {
	LogLevel    string
	AppVersion  string
	AppRevision string
}

// Configure ログ設定を実行
func Configure(config Config) {
	once.Do(func() {
		switch strings.ToUpper(config.LogLevel) {
		case "DEBUG":
			logLevel = zapcore.DebugLevel
		case "INFO":
			logLevel = zapcore.InfoLevel
		case "WARN":
			logLevel = zapcore.WarnLevel
		case "ERROR":
			logLevel = zapcore.ErrorLevel
		case "FATAL":
			logLevel = zapcore.FatalLevel
		default:
			logLevel = zapcore.InfoLevel
		}

		if len(config.AppVersion) > 0 {
			appVersion = config.AppVersion
		}
		if len(config.AppRevision) > 0 {
			appRevision = config.AppRevision
		}

		zapConfig := defaultZapConfig()
		logger, _ := zapConfig.Build()
		fields := zap.Fields(defaultZapFields()...)
		localLogger = logger.WithOptions(fields).Sugar()
	})
}

// NewLogger loggerの初期化
func NewLogger() *zap.SugaredLogger {
	if localLogger == nil {
		defaultConfig := Config{
			LogLevel:    "INFO",
			AppVersion:  "NotSetting",
			AppRevision: "NotSetting",
		}
		Configure(defaultConfig)
	}
	return localLogger
}

func defaultZapConfig() zap.Config {
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(logLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    cloudLoggingLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func defaultZapFields() []zap.Field {
	return []zap.Field{
		zap.String(versionKey, appVersion),
		zap.String(revisionKey, appRevision),
	}
}

func cloudLoggingLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("DEBUG")
	case zapcore.InfoLevel:
		enc.AppendString("INFO")
	case zapcore.WarnLevel:
		enc.AppendString("WARNING")
	case zapcore.DPanicLevel:
		enc.AppendString("ERROR")
	case zapcore.ErrorLevel:
		enc.AppendString("ERROR")
	case zapcore.FatalLevel:
		enc.AppendString("ERROR")
	default:
		enc.AppendString("ERROR")
	}
}
