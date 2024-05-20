package utility

import (
	"github.com/gofiber/fiber/v2"
)

func (u *UtilityMiddlewareService) Paginator(c *fiber.Ctx) error {
	return u.paginator(c)
}
