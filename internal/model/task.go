package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Status   Status   `json:"status"`
	Priority Priority `json:"priority"`

	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`

	Payload map[string]any `json:"payload"`
	Result  string         `json:"result,omitempty"`
	Error   string         `json:"error,omitempty"`

	Tries    int `json:"tries"`
	MaxTries int `json:"max_tries"`
}

func (t *Task) IsTerminal() bool {
	return t.Status == StatusCompleted || (t.Status == StatusFailed && t.Tries >= t.MaxTries)
}

func (t *Task) Duration() time.Duration {
	if t.StartedAt == nil || t.CompletedAt == nil {
		return 0
	}
	duration := t.CompletedAt.Sub(*t.StartedAt)
	return duration
}

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
