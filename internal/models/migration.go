package models

// Migration represents a migration script, which is a file that contains SQL.
type Migration struct {
	Name string
	Path string
}
