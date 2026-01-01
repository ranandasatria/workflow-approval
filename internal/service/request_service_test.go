package service

import (
	"context"
	"errors"
	"testing"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestService_Create_Validation(t *testing.T) {
	mockReqRepo := new(MockRequestRepo)
	mockWfRepo := new(MockWorkflowRepo)

	svc := NewRequestService(mockReqRepo, mockWfRepo)

	t.Run("Gagal jika amount nol atau negatif", func(t *testing.T) {
		req := &model.Request{
			Amount:     0,
			WorkflowID: uuid.New(),
		}

		err := svc.Create(context.Background(), req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount")
	})

	t.Run("Gagal jika Workflow tidak ditemukan", func(t *testing.T) {
		wfID := uuid.New()
		req := &model.Request{
			Amount:     5000,
			WorkflowID: wfID,
		}

		mockWfRepo.On("FindByID", mock.Anything, wfID).Return(nil, errors.New("not found"))

		err := svc.Create(context.Background(), req)

		assert.Error(t, err)
		assert.Equal(t, "workflow tidak valid", err.Error())
	})
}