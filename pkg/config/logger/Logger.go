package logger

import (
	"fmt"
	"github.com/dbacilio88/go/pkg/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func LogConfiguration(level string) (*zap.Logger, error) {
	var logger zap.AtomicLevel
	switch level {
	case "debug":
		logger = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		logger = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		logger = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		logger = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		logger = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		logger = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		//case "disabled":		logger = zap.NewAtomicLevelAt(zapcore.DPanicLevel)
		//case "deprecated":		logger = zap.NewAtomicLevelAt(zapcore.DPanicLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		MessageKey: "msg",
		//CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	filename := fmt.Sprintf("/logs/%s-%s.log", config.Config.Server.Name, config.Config.Server.Environment)

	logLumberjack := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
		MaxAge:     30,
	}
	defaultLog := logger

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(logLumberjack), defaultLog).With(
			[]zap.Field{
				zap.String("app", config.Config.Server.Name),
				zap.String("env", config.Config.Server.Environment),
			},
		),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), defaultLog).With(
			[]zap.Field{
				zap.String("app", config.Config.Server.Name),
				zap.String("env", config.Config.Server.Environment),
			}),
	)
	instance := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return instance, nil
}
