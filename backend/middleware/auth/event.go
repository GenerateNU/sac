package auth

import (
	"slices"
	"time"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the host club of this event to make this request, skips check if super user
func (m *AuthMiddlewareHandler) EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
	strUser, err := sothic.GetFromSession("user", c)
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:    "redirect",
			Value:   c.OriginalURL(),
			Expires: time.Now().Add(5 * time.Minute),
			// MARK: secure should be true in prod
			// use go build tags to do this
			HTTPOnly: true,
		})

		return c.Redirect("/api/v1/auth/login")
	}

	user := models.UnmarshalUser(strUser)

	eventUUID, err := extractor(c)
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
