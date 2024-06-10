package auth

import (
	"slices"
	"time"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareHandler) ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
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

	if user.Role == models.Super {
		locals.SetUser(c, user)
		return c.Next()
	}

	clubUUID, err := extractor(c)
	if err != nil {
		return err
	}

	clubAdmin, err := clubs.GetAdminIDs(m.db, *clubUUID)
	if err != nil {
		return err
	}

	if slices.Contains(clubAdmin, user.ID) {
		locals.SetUser(c, user)
		return c.Next()
	}

	return utilities.Forbidden()
}
