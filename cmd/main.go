package main

import (
	"gotor/cfg"
	"gotor/pkg/logging"

	"go.uber.org/zap"
)

func main() {
	// init logger
	logger := logging.Logger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("syncing logger", zap.Error(err))
		}
	}(logger)

	// parse config
	config, err := cfg.NewConfig()
	if err != nil {
		logger.Fatal("parsing config", zap.Error(err))
	}

	// do something
	_ = config

	// exit
	logger.Info("exiting")
}
