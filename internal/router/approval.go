package router

import (
	"workflow-approval/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func approvalRouter(r fiber.Router, appH *handler.ApprovalHandler, rejH *handler.RejectHandler) {
	r.Post("/:id/approve", appH.Approve)
	r.Post("/:id/reject", rejH.Reject)
}