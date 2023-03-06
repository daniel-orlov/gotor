package migrator

import (
	"context"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"gotor/internal/models"
)

type Service struct {
	logger   *zap.Logger
	migrator *migrate.Migrate
}

func NewService(logger *zap.Logger, migrator *migrate.Migrate) *Service {
	return &Service{logger: logger, migrator: migrator}
}

func (s *Service) MigrateUp(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating up", zap.Int("migrations_count", len(migrations)))

	return nil
}

func (s *Service) MigrateDown(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating down", zap.Int("migrations_count", len(migrations)))

	return nil
}

func (s *Service) migrateUp(ctx context.Context, migration *models.Migration) error {
	s.logger.Info("running a migration up", zap.Any("migration", migration))

	if migration.IsLatest() {
		s.logger.Info("migrating table up to the latest version", zap.Any("table_name", migration.TableName))

		if err := s.migrator.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				s.logger.Info(
					"no migrations to apply, table is up-to-date",
					zap.Any("table_name", migration.TableName),
				)

				return nil
			}

			return errors.Wrap(err, "migrating table up to the latest version")
		}
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

	if err := s.migrator.Steps(steps); err != nil {
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

	if migration.IsLatest() {
		s.logger.Info("migrating table down to the latest version", zap.Any("table_name", migration.TableName))

		if err := s.migrator.Down(); err != nil {
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

	if err := s.migrator.Steps(-steps); err != nil {
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
