package service

import (
	"context"
	"errors"
	"time"
	"workflow-approval/internal/model"
	"workflow-approval/internal/repository"

	"github.com/google/uuid"
)

type RequestService interface {
	Create(ctx context.Context, req *model.Request) error
	GetAll(ctx context.Context) ([]model.Request, error)
}

type requestService struct {
	repo   repository.RequestRepository
	wfRepo repository.WorkflowRepository
}

func NewRequestService(r repository.RequestRepository, wfRepo repository.WorkflowRepository) RequestService {
	return &requestService{
		repo:   r,
		wfRepo: wfRepo,
	}
}

func (s *requestService) Create(ctx context.Context, req *model.Request) error {
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	_, err := s.wfRepo.FindByID(ctx, req.WorkflowID)
	if err != nil {
		return errors.New("workflow tidak valid")
	}

	req.ID = uuid.New()
	req.Status = "PENDING"
	req.CurrentStep = 1
	req.CreatedAt = time.Now()
	return s.repo.Create(ctx, req)
}

func (s *requestService) GetAll(ctx context.Context) ([]model.Request, error) {
	return s.repo.FindAll(ctx)
}
