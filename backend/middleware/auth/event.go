package auth

import (
	"net/http"
	"slices"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the host club of this event to make this request, skips check if super user
func (m *AuthMiddlewareHandler) EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
	strUser, err := sothic.GetFromSession("microsoftonline", c)
	if err != nil {
		c.Set("redirect", "/api/v1/auth/login")
		return c.SendStatus(http.StatusUnauthorized)
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
		locals.SetUser(c, user)
		return c.Next()
	}

	return utilities.Forbidden()
}
