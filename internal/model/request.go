package model

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WorkflowID  uuid.UUID `gorm:"type:uuid" json:"workflow_id"`
	CurrentStep int       `json:"current_step"`
	Status      string    `gorm:"default:PENDING" json:"status"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}