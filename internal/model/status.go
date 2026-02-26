package model

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	// StatusCanceled   Status = "canceled"
	// StatusScheduled  Status = "scheduled"
)
