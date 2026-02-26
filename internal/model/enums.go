package model

import "fmt"

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	// StatusCanceled   Status = "canceled"
	// StatusScheduled  Status = "scheduled"
)

type Priority int

const (
	PriorityHigh Priority = iota + 1
	PriorityMed
	PriorityLow
)

func (p Priority) String() string {
	switch p {
	case PriorityHigh:
		return "high"
	case PriorityMed:
		return "medium"
	case PriorityLow:
		return "low"
	default:
		return fmt.Sprintf("unknown: %d", p) // should never be this
	}
}

func ParsePriority(p string) (Priority, error) {
	switch p {
	case "high":
		return PriorityHigh, nil
	case "med":
		return PriorityMed, nil
	case "medium":
		return PriorityMed, nil
	case "low":
		return PriorityLow, nil
	default:
		return 0, fmt.Errorf("malformed priority string: %q", p)
	}
}
