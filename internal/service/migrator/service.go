package migrator

import (
	"context"
	"fmt"
	"gotor/cfg"
	"gotor/internal/models"
	"strconv"

	"github.com/golang-migrate/migrate/v4/database"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file" // for migrate.NewWithDatabaseInstance

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	config *cfg.Config
	driver database.Driver
}

func New(logger *zap.Logger, config *cfg.Config, driver database.Driver) *Service {
	return &Service{logger: logger, config: config, driver: driver}
}

func (s *Service) createMigrator(schemaPath string) (*migrate.Migrate, error) {
	source := fmt.Sprintf("file://%s/%s/", s.config.MigrationsDir, schemaPath)

	m, err := migrate.NewWithDatabaseInstance(
		source,
		s.config.Database.Name,
		s.driver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "initializing migrator with database instance")
	}

	return m, nil
}

func (s *Service) MigrateUp(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating up", zap.Int("migrations_count", len(migrations)))

	for _, migration := range migrations {
		mInLoop := migration

		if err := s.migrateUp(ctx, &mInLoop); err != nil {
			return errors.Wrap(err, "migrating up")
		}
	}

	s.logger.Info("migrations are completed")

	return nil
}

func (s *Service) MigrateDown(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating down", zap.Int("migrations_count", len(migrations)))

	for _, migration := range migrations {
		mInLoop := migration

		if err := s.migrateDown(ctx, &mInLoop); err != nil {
			return errors.Wrap(err, "migrating down")
		}
	}

	s.logger.Info("migrations are completed")

	return nil
}

func (s *Service) migrateUp(ctx context.Context, migration *models.Migration) error {
	s.logger.Info("running a migration up", zap.Any("migration", migration))

	migrator, err := s.createMigrator(migration.Path)
	if err != nil {
		return errors.Wrap(err, "creating a migrator")
	}

	defer func(migrator *migrate.Migrate) {
		err, _ = migrator.Close()
		if err != nil {
			s.logger.Error("closing a migrator", zap.Error(err))
		}
	}(migrator)

	go func() {
		<-ctx.Done()
		s.logger.Info("stopping a migrator")
		migrator.GracefulStop <- true
	}()

	if migration.IsLatest() {
		s.logger.Info("migrating table up to the latest version", zap.Any("table_name", migration.TableName))

		if err = migrator.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				s.logger.Info(
					"no migrations to apply, table is up-to-date",
					zap.Any("table_name", migration.TableName),
				)

				return nil
			}

			return errors.Wrap(err, "migrating table up to the latest version")
		}

		return nil
	}

	s.logger.Info(
		"migrating table up to a specific version",
		zap.String("table_name", migration.TableName),
		zap.String("version", migration.Version),
	)

	// Convert the version to an int using Atoi.
	steps, err := strconv.Atoi(migration.Version)
	if err != nil {
		return errors.Wrap(err, "converting migration version to an int")
	}

	if err = migrator.Steps(steps); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			s.logger.Info(
				"no migrations to apply, table is up-to-date",
				zap.String("table_name", migration.TableName),
				zap.String("version", migration.Version),
			)

			return nil
		}

		return errors.Wrap(err, "migrating table up to a specific version")
	}

	return nil
}

func (s *Service) migrateDown(ctx context.Context, migration *models.Migration) error {
	s.logger.Info("running a migration down", zap.Any("migration", migration))

	migrator, err := s.createMigrator(migration.Path)
	if err != nil {
		return errors.Wrap(err, "creating a migrator")
	}

	defer func(migrator *migrate.Migrate) {
		err, _ = migrator.Close()
		if err != nil {
			s.logger.Error("closing a migrator", zap.Error(err))
		}
	}(migrator)

	go func() {
		<-ctx.Done()
		s.logger.Info("stopping a migrator")
		migrator.GracefulStop <- true
	}()

	if migration.IsLatest() {
		s.logger.Info("migrating table down to the latest version", zap.Any("table_name", migration.TableName))

		if err = migrator.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				s.logger.Info(
					"no migrations to apply, table is up-to-date",
					zap.Any("table_name", migration.TableName),
				)

				return nil
			}

			return errors.Wrap(err, "migrating table down to the latest version")
		}
	}

	s.logger.Info(
		"migrating table down to a specific version",
		zap.String("table_name", migration.TableName),
		zap.String("version", migration.Version),
	)

	// Convert the version to an int using Atoi.
	steps, err := strconv.Atoi(migration.Version)
	if err != nil {
		return errors.Wrap(err, "converting migration version to an int")
	}

	if err = migrator.Steps(-steps); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			s.logger.Info(
				"no migrations to apply, table is up-to-date",
				zap.String("table_name", migration.TableName),
				zap.String("version", migration.Version),
			)

			return nil
		}

		return errors.Wrap(err, "migrating table down to a specific version")
	}

	return nil
}
