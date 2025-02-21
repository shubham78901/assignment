package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// InitLogger initializes the Zap logger.
func InitLogger() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// GetLogger returns the Zap logger.
func GetLogger() *zap.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

// GetSugaredLogger returns a SugaredLogger for formatted logs.
func GetSugaredLogger() *zap.SugaredLogger {
	return GetLogger().Sugar()
}
