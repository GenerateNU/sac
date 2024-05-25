package auth

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := m.Authenticate(c); err != nil {
			return utilities.Unauthorized()
		}

		if m.IsSuper(c) {
			return c.Next()
		}

		idAsUUID, err := utilities.ValidateID(c.Params("userID"))
		if err != nil {
			return err
		}

		userID, err := auth.UserIDFrom(c)
		if err != nil {
			return err
		}

		if idAsUUID == userID {
			return c.Next()
		}

		return utilities.Forbidden()
	}(c)
}
