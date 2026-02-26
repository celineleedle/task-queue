package model

import "fmt"

type CreateTaskRequest struct {
	Type       string         `json:"type"`
	Priority   string         `json:"priority"`
	Payload    map[string]any `json:"payload"`
	MaxRetries int            `json:"max_retries"`
}

func (c *CreateTaskRequest) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("task type can not be empty")
	}

	_, err := ParsePriority(c.Priority)
	if err != nil {
		return fmt.Errorf("error parsing priority string: %w", err)
	}

	if c.MaxRetries < 0 {
		return fmt.Errorf("max retries must be a positive int: %d", c.MaxRetries)
	}

	return nil
}
