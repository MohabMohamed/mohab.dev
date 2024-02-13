package util

import (
	"log"

	"github.com/MohabMohamed/mohab.dev/src/config"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	var logger *zap.Logger
	var err error
	if config.GetEnv("ENV", "production") == "development" {

		logger, err = zap.NewDevelopment()
	} else {

		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalf("can't initialize zap Logger: %v", err)
	}
	Logger = logger.Sugar()
}

func TerminateLogger() {
	Logger.Sync()
}
