package model

import (
	"fmt"
	"testing"
)

func TestPriorityToString(t *testing.T) {
	tests := []struct {
		name  string
		input Priority
		want  string
	}{
		{"test_high", PriorityHigh, "high"},
		{"test_med", PriorityMed, "medium"},
		{"test_low", PriorityLow, "low"},
		{"test_unknown_0", 0, fmt.Sprintf("unknown: %d", 0)},
		{"test_unknown_99", 99, fmt.Sprintf("unknown: %d", 99)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.input.String()
			if res != tt.want {
				t.Errorf("Expected %q, got %q", tt.want, res)
			}
		})
	}
}

func TestParsePriority(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Priority
		wantErr bool
	}{
		{"test_high", "high", PriorityHigh, false},
		{"test_med", "med", PriorityMed, false},
		{"test_medium", "medium", PriorityMed, false},
		{"test_low", "low", PriorityLow, false},
		{"test_bad_input", "bad-input", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
