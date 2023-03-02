package main

import (
	"go.uber.org/zap"

	"gotor/cfg"
	"gotor/pkg/logging"
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
