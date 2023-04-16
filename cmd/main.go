package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"

	"gotor/cfg"
	"gotor/cli"
	"gotor/internal/service/migrator"
	"gotor/pkg/logging"

	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"
)

var k = koanf.New("/") //nolint:gochecknoglobals

func main() {
	// Parse config.
	config, err := cfg.NewConfig()
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	// Init logger.
	logger := logging.Logger(config.Logging.Format, config.Logging.Level)
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			logger.Fatal("syncing logger", zap.Error(err))
		}
	}(logger)

	logger.Info("starting migrator...")

	// Init Migrator.
	migrateInstance, err := migrate.New("file://"+config.MigrationsDir, config.Database.DSN)
	if err != nil {
		logger.Fatal("initializing migrator", zap.Error(err))
	}

	// Init services.
	migratorSvc := migrator.NewService(logger, config, migrateInstance)

	// Init cli.
	gotorCLI := cli.New(logger, migratorSvc)

	// Init context.
	ctx, cancel := initContext()
	defer cancel()

	// Run CLI.
	err = gotorCLI.Run(ctx, k)
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
