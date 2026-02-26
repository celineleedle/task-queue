package model

import (
	"testing"
)

func TestValidateCreateTaskDto(t *testing.T) {
	dtos := []*CreateTaskDto{
		&CreateTaskDto{
			Type:     "",
			Priority: "high",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "high",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "medium",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "med",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "low",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "unknown",
			Payload:  nil,
			MaxTries: 10,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "unknown",
			Payload:  nil,
			MaxTries: 0,
		},
		&CreateTaskDto{
			Type:     "type",
			Priority: "med",
			Payload:  nil,
			MaxTries: 0,
		},
		&CreateTaskDto{
			Type:     "",
			Priority: "unknown",
			Payload:  nil,
			MaxTries: 0,
		},
	}
	names := []string{
		"test_type_empty",
		"test_priority_high",
		"test_priority_medium",
		"test_priority_med",
		"test_priority_low",
		"test_priority_unknown",
		"test_priority_unknown_maxtries_0",
		"test_maxtries_0",
		"text_type_empty_priority_unknown_maxtries_0",
	}
	wantErrs := []bool{true, false, false, false, false, true, true, true, true}

	tests := make([]struct {
		name    string
		input   *CreateTaskDto
		wantErr bool
	}, 0, 9)

	for i := range 9 {
		tests = append(tests, struct {
			name    string
			input   *CreateTaskDto
			wantErr bool
		}{
			name:    names[i],
			input:   dtos[i],
			wantErr: wantErrs[i],
		})
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
