package model

import "fmt"

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

func ParsePriority(priority string) (Priority, error) {
	switch priority {
	case "high":
		return PriorityHigh, nil
	case "med":
		return PriorityMed, nil
	case "medium":
		return PriorityMed, nil
	case "low":
		return PriorityLow, nil
	default:
		return 0, fmt.Errorf("malformed priority string: %q", priority)
	}
}
