package handler

import (
	"workflow-approval/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RejectHandler struct {
	svc service.ApprovalService
}

func NewRejectHandler(svc service.ApprovalService) *RejectHandler {
	return &RejectHandler{svc}
}

func (h *RejectHandler) Reject(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   "ID tidak valid",
		})
	}

	if err := h.svc.Reject(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Request rejected successfully",
		"error":   nil,
	})
}