package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the host club of this event to make this request, skips check if super user
func (m *AuthMiddlewareService) EventAuthorizeById(c *fiber.Ctx) error {
	return utilities.Make(func(c *fiber.Ctx) error {
		if err := m.Authenticate(c); err != nil {
			return utilities.Unauthorized()
		}

		if m.IsSuper(c) {
			return c.Next()
		}

		eventUUID, err := utilities.ValidateID(c.Params("eventID"))
		if err != nil {
			return err
		}

		claims, err := auth.From(c)
		if err != nil {
			return err
		}

		issuerUUID, err := utilities.ValidateID(claims.Issuer)
		if err != nil {
			return err
		}

		eventHostAdmin, err := events.GetEventHostAdminIDs(m.DB, *eventUUID)
		if err != nil {
			return err
		}

		if slices.Contains(eventHostAdmin, *issuerUUID) {
			return c.Next()
		}

		return utilities.Forbidden()
	})(c)
}
