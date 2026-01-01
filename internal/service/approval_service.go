package service

import (
	"context"
	"errors"
	"workflow-approval/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApprovalService interface {
	Approve(ctx context.Context, id uuid.UUID) error
	Reject(ctx context.Context, id uuid.UUID) error
}

type approvalService struct {
	reqRepo  repository.RequestRepository
	stepRepo repository.StepRepository
	db       *gorm.DB
}

func NewApprovalService(r repository.RequestRepository, s repository.StepRepository, db *gorm.DB) ApprovalService {
	return &approvalService{r, s, db}
}

func (s *approvalService) Approve(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		req, err := s.reqRepo.FindByIDForUpdate(ctx, tx, id)
		if err != nil {
			return errors.New("request tidak ditemukan")
		}
		if req.Status != "PENDING" {
			return errors.New("approval tidak bisa dilakukan, status bukan PENDING")
		}

		currentStep, err := s.stepRepo.GetByWorkflowAndLevel(tx, req.WorkflowID, req.CurrentStep)
		if err != nil {
			return errors.New("tahapan workflow tidak ditemukan")
		}

		if req.Amount >= currentStep.MinAmount {
			nextStep, err := s.stepRepo.GetByWorkflowAndLevel(tx, req.WorkflowID, req.CurrentStep+1)
			if err == nil {
				req.CurrentStep = nextStep.Level
			} else {
				req.Status = "APPROVED"
			}
		} else {
			req.Status = "APPROVED"
		}

		return s.reqRepo.Update(ctx, tx, req)
	})
}
