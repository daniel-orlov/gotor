package cfg

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config is the main configuration of the application.
// It is populated from environment variables and defaults.
type Config struct {
	Logging struct {
		// LogLevel is the log level to use.
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
		// LogFormat is the log format to use.
		Format string `envconfig:"LOG_FORMAT" default:"console"`
	}
	Database struct {
		// DSN is the database connection string.
		DSN        string `envconfig:"DATABASE_DSN" default:"postgresql://postgres-test:postgres-test@localhost:5432/postgres-test?sslmode=disable"`
		Name       string `envconfig:"DATABASE_NAME" default:"postgres-test"`
		SchemaName string `envconfig:"DATABASE_SCHEMA_NAME" default:"public"`
		Driver     string `envconfig:"DATABASE_DRIVER" default:"postgres"`
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
