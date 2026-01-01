package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *approvalService) Reject(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		req, err := s.reqRepo.FindByIDForUpdate(ctx, tx, id)
		if err != nil {
			return errors.New("request tidak ditemukan")
		}

		if req.Status != "PENDING" {
			return errors.New("hanya request berstatus PENDING yang dapat di-reject")
		}

		req.Status = "REJECTED"

		return s.reqRepo.Update(ctx, tx, req)
	})
}
