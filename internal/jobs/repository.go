package jobs

import "context"

type Repository interface {
	Create(ctx context.Context, job *Job) error
	List(ctx context.Context) ([]Job, error)
}
