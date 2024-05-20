package utility

import (
	"errors"
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func (u *UtilityMiddlewareService) Limiter(rate int, expiration time.Duration) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        rate,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s-%s", c.IP(), c.Path())
		},
		LimitReached: func(c *fiber.Ctx) error {
			return utilities.NewAPIError(fiber.StatusTooManyRequests, errors.New("too many requests"))
		},
	})
}
