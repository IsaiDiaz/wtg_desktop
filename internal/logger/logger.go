package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	Init(true)
}

func Init(isDebug bool) {
	var logger *zap.Logger
	var err error

	if isDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	log = logger.Sugar()
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Sync() {
	log.Sync()
}
