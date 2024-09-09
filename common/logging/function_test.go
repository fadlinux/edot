package logging

import (
	"errors"
	"testing"
)

func TestLogInfo(t *testing.T) {
	Initialize()
	testInput := make(map[string]string)
	testInput["test"] = "testing"
	testInput["id"] = "123"
	type args struct {
		message string
		input   map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1. normal test",
			args: args{
				message: "test info",
				input:   testInput,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LogInfo(tt.args.message, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("LogInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogError(t *testing.T) {
	Initialize()
	type args struct {
		message string
		errMsg  error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1. normal test",
			args: args{
				message: "test error",
				errMsg:  errors.New("test error here"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LogError(tt.args.message, tt.args.errMsg); (err != nil) != tt.wantErr {
				t.Errorf("LogError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
