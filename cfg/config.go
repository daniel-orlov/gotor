package cfg

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config is the main configuration of the application.
// It is populated from environment variables and defaults.
type Config struct {
	// LogLevel is the log level to use.
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	Database struct {
		// DSN is the database connection string.
		DSN string `envconfig:"DATABASE_DSN" required:"true"`
	}
	// MigrationsDir is the directory where migrations are stored.
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"migrations"`
}

// NewConfig returns a new Config instance, populated with environment variables and defaults.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "processing environment variables")
	}

	return cfg, nil
}
