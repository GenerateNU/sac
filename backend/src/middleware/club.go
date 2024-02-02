package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *MiddlewareService) ClubAuthorizeById(c *fiber.Ctx) error {
	clubUUID, err := utilities.ValidateID(c.Params("id"))
	if err != nil {
		return errors.FailedToParseUUID.FiberError(c)
	}

	token, tokenErr := auth.ParseAccessToken(c.Cookies("access_token"))
	if tokenErr != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}

	issuerUUID, issueErr := utilities.ValidateID(claims.Issuer)
	if issueErr != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	// use club_id to get the list of admin for a certain club
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
