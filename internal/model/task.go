package model

import "time"

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
