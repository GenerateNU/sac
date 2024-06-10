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
	limiter   store.Storer
}

func New(paginator fiber.Handler, limiter store.Storer) Service {
	return &Handler{
		paginator: paginator,
		limiter:   limiter,
	}
}
