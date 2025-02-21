// internal/logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func InitLogger() {
	logger, _ := zap.NewProduction()
	log = logger
}

func GetLogger() *zap.Logger {
	return log
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}
