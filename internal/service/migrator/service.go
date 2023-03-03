package migrator

import (
	"context"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) Migrate(ctx context.Context) error {
	s.logger.Info("migrating")
	return nil
}
