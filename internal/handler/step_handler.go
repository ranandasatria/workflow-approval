package handler

import (
	"workflow-approval/internal/model"
	"workflow-approval/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StepHandler struct {
	svc service.StepService
}

func NewStepHandler(svc service.StepService) *StepHandler {
	return &StepHandler{svc}
}

func (h *StepHandler) Create(c *fiber.Ctx) error {
	workflowID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Workflow ID tidak valid",
		})
	}

	var step model.WorkflowStep
	if err := c.BodyParser(&step); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Payload tidak valid",
		})
	}

	step.WorkflowID = workflowID
	if err := h.svc.CreateStep(c.Context(), &step); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    step,
		"error":   nil,
	})
}

func (h *StepHandler) ListByWorkflow(c *fiber.Ctx) error {
	workflowID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "Workflow ID tidak valid",
		})
	}

	res, err := h.svc.GetByWorkflow(c.Context(), workflowID)
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