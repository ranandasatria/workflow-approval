package service

import (
	"context"
	"testing"
	"workflow-approval/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestApprove_Success(t *testing.T) {
	dbSql, mockSql, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbSql.Close()

	dialector := postgres.New(postgres.Config{Conn: dbSql})
	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	mockSql.ExpectBegin()
	mockSql.ExpectCommit()

	mockReqRepo := new(MockRequestRepo)
	mockStepRepo := new(MockStepRepo)
	svc := NewApprovalService(mockReqRepo, mockStepRepo, db)

	reqID := uuid.New()
	wfID := uuid.New()
	req := &model.Request{
		ID:          reqID,
		WorkflowID:  wfID,
		Amount:      1000,
		CurrentStep: 1,
		Status:      "PENDING",
	}
	step := &model.WorkflowStep{
		WorkflowID: wfID,
		Level:      1,
		MinAmount:  5000, 
	}

	mockReqRepo.On("FindByIDForUpdate", mock.Anything, mock.Anything, reqID).Return(req, nil)
	mockStepRepo.On("GetByWorkflowAndLevel", mock.Anything, wfID, 1).Return(step, nil)
	mockReqRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err = svc.Approve(context.Background(), reqID)

	assert.NoError(t, err)
	assert.Equal(t, "APPROVED", req.Status)
	
	mockReqRepo.AssertExpectations(t)
	mockStepRepo.AssertExpectations(t)
}