package main

import "testing"

func TestLogrusDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"ssss"},
		{"fsfsdfs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogrusDemo(tt.name)
		})
	}
}
