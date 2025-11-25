package jobs

import "context"

type Repository interface {
	Create(ctx context.Context, job *Job) error
}
