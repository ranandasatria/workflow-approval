package service

import (
	"context"
	"errors"
	"workflow-approval/internal/model"
	"workflow-approval/internal/repository"
	"github.com/google/uuid"
	"time"
)

type WorkflowService interface {
	Create(ctx context.Context, name string) (*model.Workflow, error)
	GetAll(ctx context.Context) ([]model.Workflow, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Workflow, error)
}

type workflowService struct {
	repo repository.WorkflowRepository
}

func NewWorkflowService(r repository.WorkflowRepository) WorkflowService {
	return &workflowService{r}
}

func (s *workflowService) Create(ctx context.Context, name string) (*model.Workflow, error) {
	if name == "" {
		return nil, errors.New("workflow name wajib diisi")
	}
	wf := &model.Workflow{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	return wf, s.repo.Create(ctx, wf)
}

func (s *workflowService) GetAll(ctx context.Context) ([]model.Workflow, error) {
	return s.repo.FindAll(ctx)
}

func (s *workflowService) GetByID(ctx context.Context, id uuid.UUID) (*model.Workflow, error) {
	return s.repo.FindByID(ctx, id)
}