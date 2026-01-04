package jobs

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrJobNotFound = errors.New("job not found")

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

func (s *Service) ListJobs(ctx context.Context) ([]Job, error) {
	return s.repo.List(ctx)
}

func (s *Service) CancelJob(ctx context.Context, id string) error {
	return s.repo.Cancel(ctx, id)
}
