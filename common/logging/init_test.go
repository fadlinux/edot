package logging

import "testing"

func TestInitialize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1. normal test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize()
		})
	}
}
