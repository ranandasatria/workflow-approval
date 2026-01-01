package router

import (
	"workflow-approval/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Request  *handler.RequestHandler
	Approval *handler.ApprovalHandler
	Reject   *handler.RejectHandler
	Workflow *handler.WorkflowHandler
	Step     *handler.StepHandler
}

func SetupRoutes(app *fiber.App, h Handlers) {
	api := app.Group("/api")

	workflowRouter(api.Group("/workflows"), h.Workflow)
	
	stepRouter(api.Group("/workflows/:id/steps"), h.Step)

	requestRouter(api.Group("/requests"), h.Request)
	approvalRouter(api.Group("/requests"), h.Approval, h.Reject)
}