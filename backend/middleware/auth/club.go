package auth

import (
	"slices"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareHandler) ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
	if err := m.Authenticate(c); err != nil {
		return utilities.Unauthorized()
	}

	if m.IsSuper(c) {
		return c.Next()
	}

	clubUUID, err := extractor(c)
	if err != nil {
		return err
	}

	user, err := locals.UserFrom(c)
	if err != nil {
		return err
	}

	clubAdmin, err := clubs.GetAdminIDs(m.db, *clubUUID)
	if err != nil {
		return err
	}

	if slices.Contains(clubAdmin, user.ID) {
		return c.Next()
	}

	return utilities.Forbidden()
}
