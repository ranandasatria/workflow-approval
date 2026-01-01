package handler

import (
	"workflow-approval/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ApprovalHandler struct {
	svc service.ApprovalService
}

func NewApprovalHandler(svc service.ApprovalService) *ApprovalHandler {
	return &ApprovalHandler{svc}
}

func (h *ApprovalHandler) Approve(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "ID tidak valid",
		})
	}

	if err := h.svc.Approve(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Request approved successfully",
		"error":   nil,
	})
}

