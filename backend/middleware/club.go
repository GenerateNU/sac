package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareService) ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error {
	return utilities.Make(func(c *fiber.Ctx) error {
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

		claims, err := auth.From(c)
		if err != nil {
			return err
		}

		issuerUUID, err := utilities.ValidateID(claims.Issuer)
		if err != nil {
			return err
		}

		clubAdmin, err := clubs.GetAdminIDs(m.DB, *clubUUID)
		if err != nil {
			return err
		}

		if slices.Contains(clubAdmin, *issuerUUID) {
			return c.Next()
		}

		return utilities.Forbidden()
	})(c)
}
