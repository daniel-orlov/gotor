package cli

import (
	"context"

	"go.uber.org/zap"

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

func (c *CLI) Run(ctx context.Context, cmd models.Command, migrations []models.Migration) error {
	c.logger.Info("running command", zap.String("command", cmd.Type.String()))

	switch cmd.Type {
	case models.CommandMigrateUp:
		return c.migrator.MigrateUp(ctx, migrations)
	case models.CommandMigrateDown:
		return c.migrator.MigrateDown(ctx, migrations)
	default:
		return nil
	}
}
