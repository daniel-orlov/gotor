package migrator

import (
	"context"
	"gotor/internal/models"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) MigrateUp(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating up", zap.Any("migrations", migrations))

	return nil
}

func (s *Service) MigrateDown(ctx context.Context, migrations []models.Migration) error {
	s.logger.Info("migrating down", zap.Any("migrations", migrations))

	return nil
}
