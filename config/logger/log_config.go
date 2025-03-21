package logger

import (
	env "TestProject/config"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	var err error
	if env.GetEnv("ENV", "dev") == "dev" {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}
	if err != nil {
		println("logger init error")
		panic(err)
	}
}

func Stop() {
	if log != nil {
		_ = log.Sync()
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

func Logger() *zap.Logger {
	return log
}
