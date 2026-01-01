package handler

import (
	"workflow-approval/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WorkflowHandler struct {
	svc service.WorkflowService
}

func NewWorkflowHandler(svc service.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{svc}
}

func (h *WorkflowHandler) Create(c *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Payload tidak valid",
		})
	}

	res, err := h.svc.Create(c.Context(), body.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"error":   nil,
	})
}

func (h *WorkflowHandler) List(c *fiber.Ctx) error {
	res, err := h.svc.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"error":   nil,
	})
}

func (h *WorkflowHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "ID tidak valid",
		})
	}

	res, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Workflow tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    res,
		"error":   nil,
	})
}