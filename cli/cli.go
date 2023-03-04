package cli

import (
	"context"

	"go.uber.org/zap"

	"gotor/internal/models"
)

type Migrator interface {
	MigrateUp(ctx context.Context, migration *models.Migration) error
	MigrateDown(ctx context.Context, migration *models.Migration) error
}

type CLI struct {
	logger   *zap.Logger
	migrator Migrator
}
