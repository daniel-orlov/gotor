package models

import (
	"strings"

	"github.com/pkg/errors"
)

// Command is a command to be executed by the migrator.
type Command struct {
	Type CommandType
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

var (
	// ErrUnknownCommandType is returned when an unknown command type is parsed.
	ErrUnknownCommandType = errors.New("unknown command type")
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

func ParseCommandType(s string) (CommandType, error) {
	switch strings.ToLower(s) {
	case CommandMigrateUpString:
		return CommandMigrateUp, nil
	case CommandMigrateDownString:
		return CommandMigrateDown, nil
	default:
		return CommandUnknown, ErrUnknownCommandType
	}
}

func AllowedCommands() []string {
	return []string{
		CommandMigrateUpString,
		CommandMigrateDownString,
	}
}

func AllCommandsString() string {
	return strings.Join(AllowedCommands(), ", ")
}
