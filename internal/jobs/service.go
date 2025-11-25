package jobs

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type CreateJobInput struct {
	Type    string         `json:"type" binding:"required"`
	Payload map[string]any `json:"payload"`
}

func (s *Service) CreateJob(ctx context.Context, in CreateJobInput) (*Job, error) {
	job := &Job{
		ID:        uuid.New(),
		Type:      in.Type,
		Status:    StatusPending,
		Payload:   in.Payload,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, job); err != nil {
		return nil, err
	}

	return job, nil
}
