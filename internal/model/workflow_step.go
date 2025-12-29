package model

import (
	"github.com/google/uuid"
)

type WorkflowStep struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	WorkflowID uuid.UUID `gorm:"type:uuid;not null"`
	Level      int       `gorm:"not null"`
	Role      string    `gorm:"not null"`
	MinAmount  int64     `gorm:"not null"`
}
