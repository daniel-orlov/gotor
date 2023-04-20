package flags

import (
	"fmt"
	"gotor/internal/models"
	"os"

	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
)

const (
	// FlagCommand is the name of the command flag. It is expected to be one of the allowed commands.
	FlagCommand = "command"
	// FlagMigrations is the name of the migrations flag. It is a comma separated list of migrations.
	FlagMigrations = "migrations"
)

// Parse parses the command line flags and returns the command to run.
func Parse(k *koanf.Koanf) error {
	flagSet := flag.NewFlagSet("commands", flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Println(flagSet.FlagUsages())
		os.Exit(0)
	}

	flagSet.StringSlice(
		FlagMigrations,
		[]string{""},
		"Comma separated list of migrations to run. "+
			"Format <table_name>:<migration_version>. "+
			"If empty, all migrations will be run to the latest version (only applicable to upwards direction).",
	)

	flagSet.String(
		FlagCommand,
		models.CommandMigrateUpString,
		"Command to run. One of: "+models.AllCommandsString(),
	)

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "parsing flags")
	}

	if err = k.Load(posflag.Provider(flagSet, "/", k), nil); err != nil {
		return errors.Wrap(err, "loading flags")
	}

	return nil
}
