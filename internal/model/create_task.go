package model

import "fmt"

type CreateTaskRequest struct {
	Type     string         `json:"type"`
	Priority string         `json:"priority"`
	Payload  map[string]any `json:"payload"`
	MaxTries int            `json:"max_retries"`
}

func (c *CreateTaskRequest) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("task type can not be empty")
	}

	_, err := ParsePriority(c.Priority)
	if err != nil {
		return fmt.Errorf("error parsing priority string: %w", err)
	}

	if c.MaxTries < 1 {
		return fmt.Errorf("max tries must be greater than one")
	}

	return nil
}
