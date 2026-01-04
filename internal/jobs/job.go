package jobs

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusRunning  Status = "running"
	StatusFinished Status = "finished"
	StatusFailed   Status = "failed"
	StatusCanceled Status = "canceled"
)

type Job struct {
	ID        uuid.UUID      `json:"id"`
	Type      string         `json:"type"`
	Status    Status         `json:"status"`
	Payload   map[string]any `json:"payload,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
