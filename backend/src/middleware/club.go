package middleware

// SELECT * FROM user_club_members WHERE club_id = X AND membership_type = 'admin';
// Slice -> is value in slice?

import (
	"slices"
	"strconv"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func (m *MiddlewareService) ClubAuthorizeById(c *fiber.Ctx) error {
	clubId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	token, err := auth.ParseAccessToken(c.Cookies("access_token"))
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	issuerID, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	// use club_id to get the list of admin for a certain club
	clubAdmin, err := transactions.GetAdminIDs(m.DB, clubId)

	// check issuerID against the list of admin for the certain club

	if slices.Contains(clubAdmin, issuerID) {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "unauthorized",
	})
}
