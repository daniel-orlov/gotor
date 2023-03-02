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
	// Source is the source of the migrations.
	Source string `envconfig:"SOURCE" default:"./migrations"`
	// DatabaseURL is the URL of the database to connect to.
	DatabaseURL string `envconfig:"DATABASE_URL" default:""`
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
