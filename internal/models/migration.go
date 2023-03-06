package models

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Migration represents a migration script, which is a file that contains SQL.
type Migration struct {
	// TableName is the name of the table that migration is going to be applied to.
	TableName string
	// Path is the path to the migration file.
	Path string
	// Version is the version of the migration.
	Version string
	// Steps is the number of steps that the migration is going to take.
	Steps int
}

const (
	VersionLatest            = "latest"
	validMigrationPartsCount = 2
	defaultMigrationPath     = "internal/migrations"
)

func (m Migration) IsLatest() bool {
	return m.Version == VersionLatest
}

// ParseMigrationList parses a [] string into a []Migration.
// The string should be in the format <table_name>:<migration_version>,<table_name>:<migration_version>,
// where <table_name> is the name of the table that migration is going to be applied to and
// <migration_version> is the version of the migration.
func ParseMigrationList(migrations []string, direction string) ([]Migration, error) {
	if len(migrations) == 0 {
		return nil, errors.New("no migrations provided")
	}

	// Parse each migration.
	parsedMigrations := make([]Migration, 0, len(migrations))
	for _, migration := range migrations {
		parsedMigration, err := ParseMigration(migration, direction)
		if err != nil {
			return nil, errors.Wrap(err, "parsing migrations")
		}

		parsedMigrations = append(parsedMigrations, *parsedMigration)
	}

	return parsedMigrations, nil
}

// ParseMigration parses a string into a migration.
// The string should be in the format <table_name>:<migration_version>.
func ParseMigration(migration, direction string) (*Migration, error) {
	// Split the string into a table name and a version.
	parts := strings.Split(migration, ":")
	if len(parts) != validMigrationPartsCount {
		return nil, errors.New("invalid migration format, should be <table_name>:<migration_version>")
	}

	return &Migration{
		TableName: parts[0],
		Version:   parts[1],
		Path:      filepath.Join(defaultMigrationPath, parts[0], fmt.Sprintf("%s.%s.sql", parts[1], direction)),
	}, nil
}
