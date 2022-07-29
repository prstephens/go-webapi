package logger

import (
	"fmt"
	"os"

	"github.com/prstephens/go-webapi/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(config *config.Config) *zap.Logger {

	var defaultLogLevel zapcore.Level

	// setup prod or dev environment logging
	switch config.App.Environment {
	case "prod":
		defaultLogLevel = zapcore.InfoLevel
	default:
		defaultLogLevel = zapcore.DebugLevel
	}

	zapconfig := zap.NewProductionEncoderConfig()
	zapconfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(zapconfig)
	consoleEncoder := zapcore.NewConsoleEncoder(zapconfig)

	logFile, err := os.OpenFile(config.App.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("Failed to create logfile: " + err.Error())
	}

	writer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	fmt.Println("Using following envrionment:", config.App.Environment)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
