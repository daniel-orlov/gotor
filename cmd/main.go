package main

import (
	"gotor/cfg"
	"gotor/pkg/logging"

	"go.uber.org/zap"
)

func main() {
	// Init logger.
	logger := logging.Logger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("syncing logger", zap.Error(err))
		}
	}(logger)

	logger.Info("starting migrator...")

	// Parse config.
	config, err := cfg.NewConfig()
	if err != nil {
		logger.Fatal("parsing config", zap.Error(err))
	}
	_ = config

	logger.Info("config parsed successfully")

	// Init services.
	migratorSvc := migrator.NewService(logger)

	// Init cli.
	gotorCLI := cli.New(logger, migratorSvc)

	// Init context.
	ctx, cancel := initContext()
	defer cancel()

	// Run CLI.
	err = gotorCLI.Run(ctx)
	if err != nil {
		logger.Fatal("running cli", zap.Error(err))
	}

	logger.Info("exiting migrator...")
}
