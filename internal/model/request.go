package model

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	WorkflowID   uuid.UUID `gorm:"type:uuid;not null"`
	CurrentStep  int       `gorm:"not null"`
	Status       string    `gorm:"not null"`
	Amount       int64     `gorm:"not null"`
	CreatedAt    time.Time
}
