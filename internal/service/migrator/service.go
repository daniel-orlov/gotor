package migrator

import (
	"context"

	"go.uber.org/zap"

	"gotor/internal/models"
)

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) MigrateUp(ctx context.Context, migration *models.Migration) error {
	s.logger.Info("migrating up")

	return nil
}

func (s *Service) MigrateDown(ctx context.Context, migration *models.Migration) error {
	s.logger.Info("migrating down")

	return nil
}
