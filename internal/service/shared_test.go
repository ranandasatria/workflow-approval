package service

import (
	"context"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockRequestRepo struct{ mock.Mock }

func (m *MockRequestRepo) Create(ctx context.Context, req *model.Request) error {
	return m.Called(ctx, req).Error(0)
}
func (m *MockRequestRepo) FindAll(ctx context.Context) ([]model.Request, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Request), args.Error(1)
}
func (m *MockRequestRepo) FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*model.Request, error) {
	args := m.Called(ctx, tx, id)
	return args.Get(0).(*model.Request), args.Error(1)
}
func (m *MockRequestRepo) Update(ctx context.Context, tx *gorm.DB, req *model.Request) error {
	return m.Called(ctx, tx, req).Error(0)
}

type MockWorkflowRepo struct{ mock.Mock }

func (m *MockWorkflowRepo) FindByID(ctx context.Context, id uuid.UUID) (*model.Workflow, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Workflow), args.Error(1)
}
func (m *MockWorkflowRepo) Create(ctx context.Context, wf *model.Workflow) error {
	return m.Called(ctx, wf).Error(0)
}
func (m *MockWorkflowRepo) FindAll(ctx context.Context) ([]model.Workflow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Workflow), args.Error(1)
}

type MockStepRepo struct{ mock.Mock }

func (m *MockStepRepo) GetByWorkflowAndLevel(tx *gorm.DB, workflowID uuid.UUID, level int) (*model.WorkflowStep, error) {
	args := m.Called(tx, workflowID, level)
	return args.Get(0).(*model.WorkflowStep), args.Error(1)
}
func (m *MockStepRepo) Create(ctx context.Context, step *model.WorkflowStep) error {
	return m.Called(ctx, step).Error(0)
}
func (m *MockStepRepo) FindByWorkflowID(ctx context.Context, workflowID uuid.UUID) ([]model.WorkflowStep, error) {
	args := m.Called(ctx, workflowID)
	return args.Get(0).([]model.WorkflowStep), args.Error(1)
}