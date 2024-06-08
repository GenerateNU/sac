package auth

import (
	"slices"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the host club of this event to make this request, skips check if super user
func (m *AuthMiddlewareHandler) EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
	if err := m.Authenticate(c); err != nil {
		return utilities.Unauthorized()
	}

	if m.IsSuper(c) {
		return c.Next()
	}

	eventUUID, err := extractor(c)
	if err != nil {
		return err
	}

	user, err := locals.UserFrom(c)
	if err != nil {
		return err
	}

	eventHostAdmin, err := events.GetEventHostAdminIDs(m.db, *eventUUID)
	if err != nil {
		return err
	}

	if slices.Contains(eventHostAdmin, user.ID) {
		return c.Next()
	}

	return utilities.Forbidden()
}
