package storage

import (
	"context"
	"encoding/json"

	"github.com/Facundoblanco10/go-pulse-core/internal/jobs"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) Create(ctx context.Context, j *jobs.Job) error {
	model := JobModel{
		ID:     j.ID,
		Type:   j.Type,
		Status: string(j.Status),
	}

	if j.Payload != nil {
		jsonBytes, _ := json.Marshal(j.Payload)
		model.Payload = datatypes.JSON(jsonBytes)
	}

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	j.ID = model.ID
	return nil
}

func (r *JobRepository) List(ctx context.Context) ([]jobs.Job, error) {
	var result []JobModel
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}

	jobsList := make([]jobs.Job, len(result))
	for i, model := range result {
		var payload map[string]any
		if model.Payload != nil {
			_ = json.Unmarshal(model.Payload, &payload)
		}
		jobsList[i] = jobs.Job{
			ID:        model.ID,
			Type:      model.Type,
			Status:    jobs.Status(model.Status),
			Payload:   payload,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}

	return jobsList, nil
}

func (r *JobRepository) Cancel(ctx context.Context, id string) error {
	res := r.db.WithContext(ctx).Model(&JobModel{}).Where("id = ?", id).
		Update("status", string(jobs.StatusCanceled))
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return jobs.ErrJobNotFound
	}

	return nil
}
