package cli

import (
	"context"

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

func (c *CLI) Run(ctx context.Context) error {
	cmd, migrations, err := flags.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing flags")
	}

	c.logger.Info(
		"running command",
		zap.String("command", cmd.String()),
		zap.Int("migrations_count", len(migrations)),
	)

	switch *cmd {
	case models.CommandMigrateUp:
		return c.migrator.MigrateUp(ctx, migrations)
	case models.CommandMigrateDown:
		return c.migrator.MigrateDown(ctx, migrations)
	default:
		return nil
	}
}
