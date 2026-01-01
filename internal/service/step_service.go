package service

import (
	"context"
	"errors"
	"workflow-approval/internal/model"
	"workflow-approval/internal/repository"
	"github.com/google/uuid"
)

type StepService interface {
	CreateStep(ctx context.Context, step *model.WorkflowStep) error
	GetByWorkflow(ctx context.Context, workflowID uuid.UUID) ([]model.WorkflowStep, error)
}

type stepService struct {
	repo repository.StepRepository
}

func NewStepService(r repository.StepRepository) StepService {
	return &stepService{r}
}

func (s *stepService) CreateStep(ctx context.Context, step *model.WorkflowStep) error {
	existing, _ := s.repo.GetByWorkflowAndLevel(nil, step.WorkflowID, step.Level)
	if existing != nil && existing.ID != uuid.Nil {
		return errors.New("step level harus unik per workflow")
	}
	step.ID = uuid.New()
	return s.repo.Create(ctx, step)
}

func (s *stepService) GetByWorkflow(ctx context.Context, id uuid.UUID) ([]model.WorkflowStep, error) {
	return s.repo.FindByWorkflowID(ctx, id)
}