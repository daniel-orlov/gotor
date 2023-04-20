package main

import (
	"context"
	"gotor/cfg"
	"gotor/cli"
	"gotor/internal/service/migrator"
	"gotor/pkg/logging"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // for migrate.NewWithDatabaseInstance
	"github.com/jmoiron/sqlx"

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

	// Connect to database.
	db, err := sqlx.Connect(config.Database.Driver, config.Database.DSN)
	if err != nil {
		logger.Fatal("connecting to database", zap.Error(err))
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			logger.Fatal("closing database connection", zap.Error(err))
		}
	}(db)

	logger.Info("connected to database")

	// Init driver.
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{
		DatabaseName:    config.Database.Name,
		SchemaName:      config.Database.SchemaName,
		MigrationsTable: postgres.DefaultMigrationsTable,
	})
	if err != nil {
		logger.Fatal("initializing driver", zap.Error(err))
	}

	logger.Info("initialized driver")

	// Init services.
	migratorSvc := migrator.New(logger, config, driver)

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
