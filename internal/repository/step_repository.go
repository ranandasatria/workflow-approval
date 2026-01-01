package repository

import (
	"context"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stepRepo struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) StepRepository {
	return &stepRepo{db: db}
}

type StepRepository interface {
	GetByWorkflowAndLevel(tx *gorm.DB, workflowID uuid.UUID, level int) (*model.WorkflowStep, error)
	Create(ctx context.Context, step *model.WorkflowStep) error
	FindByWorkflowID(ctx context.Context, workflowID uuid.UUID) ([]model.WorkflowStep, error)
}

func (r *stepRepo) GetByWorkflowAndLevel(tx *gorm.DB, workflowID uuid.UUID, level int) (*model.WorkflowStep, error) {
	var step model.WorkflowStep
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.First(&step, "workflow_id = ? AND level = ?", workflowID, level).Error
	return &step, err
}

func (r *stepRepo) Create(ctx context.Context, step *model.WorkflowStep) error {
	return r.db.WithContext(ctx).Create(step).Error
}

func (r *stepRepo) FindByWorkflowID(ctx context.Context, workflowID uuid.UUID) ([]model.WorkflowStep, error) {
	var steps []model.WorkflowStep
	err := r.db.WithContext(ctx).Where("workflow_id = ?", workflowID).Order("level asc").Find(&steps).Error
	return steps, err
}