package router

import (
	"workflow-approval/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func requestRouter(r fiber.Router, h *handler.RequestHandler) {
	r.Post("", h.Create)
	r.Get("", h.FetchAll)
}