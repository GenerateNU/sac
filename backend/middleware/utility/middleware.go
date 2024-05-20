package utility

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type UtilityMiddlewareInterface interface {
	Paginator(c *fiber.Ctx) error
	Limiter(rate int, duration time.Duration) fiber.Handler
}

type UtilityMiddlewareService struct {
	paginator fiber.Handler
}

func New(paginator fiber.Handler) *UtilityMiddlewareService {
	return &UtilityMiddlewareService{
		paginator: paginator,
	}
}
