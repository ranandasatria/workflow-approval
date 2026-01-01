package app

import (
	"workflow-approval/internal/handler"
	"workflow-approval/internal/repository"
	"workflow-approval/internal/router"
	"workflow-approval/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	Handlers router.Handlers
}

func NewContainer(db *gorm.DB) *Container {
	reqRepo := repository.NewRequestRepository(db)
	stepRepo := repository.NewStepRepository(db)
	wfRepo := repository.NewWorkflowRepository(db)

	reqSvc := service.NewRequestService(reqRepo, wfRepo)
	appSvc := service.NewApprovalService(reqRepo, stepRepo, db)
	wfSvc := service.NewWorkflowService(wfRepo)    
	stepSvc := service.NewStepService(stepRepo)    

	handlers := router.Handlers{
		Request:  handler.NewRequestHandler(reqSvc),
		Approval: handler.NewApprovalHandler(appSvc),
		Reject:   handler.NewRejectHandler(appSvc),
		Workflow: handler.NewWorkflowHandler(wfSvc),  
		Step:     handler.NewStepHandler(stepSvc),      
	}

	return &Container{
		Handlers: handlers,
	}
}