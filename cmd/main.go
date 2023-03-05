package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gotor/cfg"
	"gotor/cli"
	"gotor/internal/service/migrator"
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

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		cancel()
	}()

	return ctx, cancel
}
