package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
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

	token, tokenErr := auth.ParseAccessToken(c.Cookies("access_token"), m.AuthSettings.AccessKey)
	if tokenErr != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	claims, ok := token.Claims.(*auth.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}

	issuerUUID, issueErr := utilities.ValidateID(claims.Issuer)
	if issueErr != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	// use clubID to get the list of admin for a certain club
	clubAdmin, clubErr := transactions.GetAdminIDs(m.DB, *clubUUID)
	if clubErr != nil {
		return err
	}

	// check issuerID against the list of admin for the certain club
	if slices.Contains(clubAdmin, *issuerUUID) {
		return c.Next()
	}

	return errors.Unauthorized.FiberError(c)
}
