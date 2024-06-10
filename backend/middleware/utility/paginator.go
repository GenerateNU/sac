package utility

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Paginator(c *fiber.Ctx) error {
	return h.paginator(c)
}
