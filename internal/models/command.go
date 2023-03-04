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

const (
	// CommandMigrateUpString is the string representation of the migrate up command.
	CommandMigrateUpString = "up"
	// CommandMigrateDownString is the string representation of the migrate down command.
	CommandMigrateDownString = "down"
)

func (c CommandType) String() string {
	switch c {
	case CommandMigrateUp:
		return CommandMigrateUpString
	case CommandMigrateDown:
		return CommandMigrateDownString
	default:
		return "unknown"
	}
}

func ParseCommandType(s string) CommandType {
	switch s {
	case CommandMigrateUpString:
		return CommandMigrateUp
	case CommandMigrateDownString:
		return CommandMigrateDown
	default:
		return CommandUnknown
	}
}
