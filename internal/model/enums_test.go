package model

import (
	"fmt"
	"testing"
)

func TestPriorityToString(t *testing.T) {
	tests := []struct {
		input Priority
		want  string
	}{
		{PriorityHigh, "high"},
		{PriorityMed, "medium"},
		{PriorityLow, "low"},
		{0, fmt.Sprintf("unknown: %d", 0)},
		{99, fmt.Sprintf("unknown: %d", 99)},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			res := tt.input.String()
			if res != tt.want {
				t.Errorf("Expected %q, got %q", tt.want, res)
			}
		})
	}
}

func TestParsePriority(t *testing.T) {
	tests := []struct {
		input   string
		want    Priority
		wantErr bool
	}{
		{"high", PriorityHigh, false},
		{"med", PriorityMed, false},
		{"medium", PriorityMed, false},
		{"low", PriorityLow, false},
		{"bad-input", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res, err := ParsePriority(tt.input)
			if err != nil && !tt.wantErr {
				t.Errorf("Got an error but did not expect one: %q", err.Error())
			}
			if res != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, res)
			}
		})
	}
}
