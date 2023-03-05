package cli

import (
	"context"

	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"gotor/cli/flags"
	"gotor/internal/models"
)

type Migrator interface {
	MigrateUp(ctx context.Context, migrations []models.Migration) error
	MigrateDown(ctx context.Context, migrations []models.Migration) error
}

type CLI struct {
	logger   *zap.Logger
	migrator Migrator
}

func New(logger *zap.Logger, migrator Migrator) *CLI {
	return &CLI{logger: logger, migrator: migrator}
}

func (c *CLI) Run(ctx context.Context, k *koanf.Koanf) error {
	err := flags.Parse(k)
	if err != nil {
		return errors.Wrap(err, "parsing flags")
	}

	rawCmd := k.String(flags.FlagCommand)
	if rawCmd == "" {
		return errors.New("command is empty")
	}

	rawMigrations := k.Strings(flags.FlagMigrations)
	if len(rawMigrations) == 0 {
		return errors.New("migrations are empty")
	}

	c.logger.Info(
		"running command",
		zap.String("command", rawCmd),
		zap.Strings("migrations", rawMigrations),
		zap.Int("migrations_count", len(rawMigrations)),
	)

	commandType, err := models.ParseCommandType(rawCmd)
	if err != nil {
		return errors.Wrap(err, "parsing command type")
	}

	migrations, err := models.ParseMigrationList(rawMigrations, commandType.String())
	if err != nil {
		return errors.Wrap(err, "parsing migrations")
	}

	switch commandType {
	case models.CommandMigrateUp:
		return c.migrator.MigrateUp(ctx, migrations)
	case models.CommandMigrateDown:
		return c.migrator.MigrateDown(ctx, migrations)
	default:
		return nil
	}
}
