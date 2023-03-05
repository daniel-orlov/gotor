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

func TestParseCommandType(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    models.CommandType
		wantErr bool
	}{
		{
			name: "Command unknown",
			args: args{
				s: "test",
			},
			want:    models.CommandUnknown,
			wantErr: true,
		},
		{
			name: "Command migrate up",
			args: args{
				s: "up",
			},
			want:    models.CommandMigrateUp,
			wantErr: false,
		},
		{
			name: "Command migrate down",
			args: args{
				s: "down",
			},
			want:    models.CommandMigrateDown,
			wantErr: false,
		},
		{
			name: "Command migrate up (uppercase)",
			args: args{
				s: "UP",
			},
			want:    models.CommandMigrateUp,
			wantErr: false,
		},
		{
			name: "Command migrate down (uppercase)",
			args: args{
				s: "DOWN",
			},
			want:    models.CommandMigrateDown,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.ParseCommandType(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
