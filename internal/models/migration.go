package models

// Migration represents a migration script, which is a file that contains SQL.
type Migration struct {
	// TableName is the name of the table that migration is going to be applied to.
	TableName string
	// Path is the path to the migration file.
	Path string
	// Version is the version of the migration.
	Version int
}
