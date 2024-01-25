package middleware

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func UserAuthorizeById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
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

	if issuerID == id {
		return c.Next()
	}
	

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "unauthorized",
	})
}