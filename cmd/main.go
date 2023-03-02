package main

import (
	"gotor/cfg"
	"gotor/pkg/logging"

	"go.uber.org/zap"
)

func main() {
	// init logger
	logger := logging.Logger()

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
