package utility

import (
	"time"

	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/gofiber/fiber/v2"
)

type UtilityMiddlewareInterface interface {
	Paginator(c *fiber.Ctx) error
	Limiter(rate int, duration time.Duration) fiber.Handler
}

type UtilityMiddlewareService struct {
	paginator fiber.Handler
	Stores    *store.Stores
}

func New(paginator fiber.Handler, stores *store.Stores) *UtilityMiddlewareService {
	return &UtilityMiddlewareService{
		paginator: paginator,
		Stores:    stores,
	}
}
