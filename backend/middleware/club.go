package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareService) ClubAuthorizeById(c *fiber.Ctx) error {
	if m.IsSuper(c) {
		return c.Next()
	}

	clubUUID, err := utilities.ValidateID(c.Params("clubID"))
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	claims, err := auth.From(c)
	if err != nil {
		return err.FiberError(c)
	}

	issuerUUID, issueErr := utilities.ValidateID(claims.Issuer)
	if issueErr != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	clubAdmin, clubErr := clubs.GetAdminIDs(m.DB, *clubUUID)
	if clubErr != nil {
		return err
	}

	if slices.Contains(clubAdmin, *issuerUUID) {
		return c.Next()
	}

	return errors.Forbidden.FiberError(c)
}
