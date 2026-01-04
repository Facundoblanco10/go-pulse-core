package storage

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JobModel struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type   string    `gorm:"type:varchar(50);not null"`
	Status string    `gorm:"type:varchar(20);not null;check:status IN ('pending','running','finished','failed','canceled')"`

	Payload   datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (jm *JobModel) BeforeCreate(tx *gorm.DB) (err error) {
	if jm.ID == uuid.Nil {
		jm.ID = uuid.New()
	}
	return nil
}
