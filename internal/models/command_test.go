package models_test

import (
	"reflect"
	"testing"

	"gotor/internal/models"
)

func TestCommandType_String(t *testing.T) {
	tests := []struct {
		name string
		c    models.CommandType
		want string
	}{
		{
			name: "Command unknown",
			c:    models.CommandUnknown,
			want: "unknown",
		},
		{
			name: "Command migrate up",
			c:    models.CommandMigrateUp,
			want: "up",
		},
		{
			name: "Command migrate down",
			c:    models.CommandMigrateDown,
			want: "down",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCommandType(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want models.CommandType
	}{
		{
			name: "Command unknown",
			args: args{s: "unknown"},
			want: models.CommandUnknown,
		},
		{
			name: "Command migrate up",
			args: args{s: models.CommandMigrateUpString},
			want: models.CommandMigrateUp,
		},
		{
			name: "Command migrate down",
			args: args{s: models.CommandMigrateDownString},
			want: models.CommandMigrateDown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.ParseCommandType(tt.args.s); got != tt.want {
				t.Errorf("ParseCommandType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllowedCommands(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Allowed commands list",
			want: []string{
				models.CommandMigrateUpString,
				models.CommandMigrateDownString,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.AllowedCommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllowedCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}
