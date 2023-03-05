package models_test

import (
	"reflect"
	"testing"

	"gotor/internal/models"
)

func TestParseMigration(t *testing.T) {
	type args struct {
		migration string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Migration
		wantErr bool
	}{
		{
			name:    "invalid migration - no colon",
			args:    args{migration: "test"},
			want:    models.Migration{},
			wantErr: true,
		},
		{
			name:    "invalid migration - too many colons",
			args:    args{migration: "test:test:test"},
			want:    models.Migration{},
			wantErr: true,
		},
		{
			name:    "valid migration",
			args:    args{migration: "test:test"},
			want:    models.Migration{TableName: "test", Version: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.ParseMigration(tt.args.migration)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMigration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMigration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMigrationList(t *testing.T) {
	type args struct {
		migrationString string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Migration
		wantErr bool
	}{
		{
			name:    "invalid migration - no migrations",
			args:    args{migrationString: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid migration - invalid migration",
			args:    args{migrationString: "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid migration - one migration",
			args:    args{migrationString: "test:test"},
			want:    []models.Migration{{TableName: "test", Version: "test"}},
			wantErr: false,
		},
		{
			name:    "valid migration - multiple migrations",
			args:    args{migrationString: "test:test,test2:test2,test3:test3"},
			want:    []models.Migration{{TableName: "test", Version: "test"}, {TableName: "test2", Version: "test2"}, {TableName: "test3", Version: "test3"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.ParseMigrationList(tt.args.migrationString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMigrationList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMigrationList() got = %v, want %v", got, tt.want)
			}
		})
	}
}