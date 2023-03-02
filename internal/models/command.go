package models

// Command is a command to be executed by the migrator.
type Command struct {
	Type      CommandType
	Arguments []string
}

// CommandType is the type of command.
type CommandType int

const (
	// CommandUnknown is an unknown command.
	CommandUnknown CommandType = iota
	// CommandMigrateUp is a command to migrate up.
	CommandMigrateUp
	// CommandMigrateDown is a command to migrate down.
	CommandMigrateDown
)
