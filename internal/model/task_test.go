package model

import (
	"testing"
)

func TestValidateTaskDto(t *testing.T) {
	tests := []struct {
		name    string
		input   *TaskDto
		wantErr bool
	}{
		{
			"test_type_empty",
			&TaskDto{
				Type:     "",
				Priority: "high",
				Payload:  nil,
				MaxTries: 10,
			},
			true,
		},
		{
			"test_priority_high",
			&TaskDto{
				Type:     "type",
				Priority: "high",
				Payload:  nil,
				MaxTries: 10,
			},
			false,
		},
		{
			"test_priority_medium",
			&TaskDto{
				Type:     "type",
				Priority: "medium",
				Payload:  nil,
				MaxTries: 10,
			},
			false,
		},
		{
			"test_priority_med",
			&TaskDto{
				Type:     "type",
				Priority: "med",
				Payload:  nil,
				MaxTries: 10,
			},
			false,
		},
		{
			"test_priority_low",
			&TaskDto{
				Type:     "type",
				Priority: "low",
				Payload:  nil,
				MaxTries: 10,
			},
			false,
		},
		{
			"test_priority_unknown",
			&TaskDto{
				Type:     "type",
				Priority: "unknown",
				Payload:  nil,
				MaxTries: 10,
			},
			true,
		},
		{
			"test_priority_unknown_maxtries_0",
			&TaskDto{
				Type:     "type",
				Priority: "unknown",
				Payload:  nil,
				MaxTries: 0,
			},
			true,
		},
		{
			"test_maxtries_0",
			&TaskDto{
				Type:     "type",
				Priority: "med",
				Payload:  nil,
				MaxTries: 0,
			},
			true,
		},
		{
			"text_type_empty_priority_unknown_maxtries_0",
			&TaskDto{
				Type:     "",
				Priority: "unknown",
				Payload:  nil,
				MaxTries: 0,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if err != nil && !tt.wantErr {
				t.Errorf("Got an error but did not expect one: %q", err.Error())
			} else if err == nil && tt.wantErr {
				t.Errorf("No error but wanted one")
			}
		})
	}
}
