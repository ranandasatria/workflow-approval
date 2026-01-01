package router

import (
	"workflow-approval/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func workflowRouter(r fiber.Router, h *handler.WorkflowHandler) {
	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/:id", h.GetByID)
}