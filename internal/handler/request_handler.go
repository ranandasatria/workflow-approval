package handler

import (
	"workflow-approval/internal/model"
	"workflow-approval/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler struct {
	svc service.RequestService
}

func NewRequestHandler(svc service.RequestService) *RequestHandler {
	return &RequestHandler{svc}
}

func (h *RequestHandler) Create(c *fiber.Ctx) error {
	var req model.Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Payload tidak valid",
		})
	}

	if err := h.svc.Create(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    req,
		"error":   nil,
	})
}

func (h *RequestHandler) FetchAll(c *fiber.Ctx) error {
	requests, err := h.svc.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    requests,
		"error":   nil,
	})
}
