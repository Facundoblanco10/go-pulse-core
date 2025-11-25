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
