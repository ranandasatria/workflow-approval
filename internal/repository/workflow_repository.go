package repository

import (
	"context"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowRepository interface {
	Create(ctx context.Context, workflow *model.Workflow) error
	FindAll(ctx context.Context) ([]model.Workflow, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Workflow, error)
}

type workflowRepo struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) WorkflowRepository {
	return &workflowRepo{db}
}

func (r *workflowRepo) Create(ctx context.Context, workflow *model.Workflow) error {
	return r.db.WithContext(ctx).Create(workflow).Error
}

func (r *workflowRepo) FindAll(ctx context.Context) ([]model.Workflow, error) {
	var workflows []model.Workflow
	err := r.db.WithContext(ctx).Order("created_at desc").Find(&workflows).Error
	return workflows, err
}

func (r *workflowRepo) FindByID(ctx context.Context, id uuid.UUID) (*model.Workflow, error) {
	var workflow model.Workflow
	err := r.db.WithContext(ctx).First(&workflow, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}
