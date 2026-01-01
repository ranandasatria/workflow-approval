package router

import (
	"workflow-approval/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func stepRouter(r fiber.Router, h *handler.StepHandler) {
	r.Post("/", h.Create)
	r.Get("/", h.ListByWorkflow)
}