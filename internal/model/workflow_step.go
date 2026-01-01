package model

import (
	"github.com/google/uuid"
)

type WorkflowStep struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WorkflowID uuid.UUID `gorm:"type:uuid" json:"workflow_id"`
	Level      int       `json:"level"`
	Actor      string    `json:"actor"`
	MinAmount  float64   `json:"min_amount"`
}