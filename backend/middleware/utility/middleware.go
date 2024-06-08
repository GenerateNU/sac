package utility

import (
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Paginator(c *fiber.Ctx) error
	Limiter(rate int, duration time.Duration) fiber.Handler
}

type Handler struct {
	paginator fiber.Handler
	limiter   store.Limiter
}

func New(paginator fiber.Handler, limiter store.Limiter) Service {
	return &Handler{
		paginator: paginator,
		limiter:   limiter,
	}
}
