package auth

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) UserAuthorizeById(c *fiber.Ctx) error {
	strUser, err := sothic.GetFromSession("user", c)
	if err != nil {
		c.Set("redirect", "/api/v1/auth/login")
		return c.SendStatus(http.StatusUnauthorized)
	}

	user := *models.UnmarshalUser(strUser)

	if user.Role == models.Super {
		locals.SetUser(c, &user)
		return c.Next()
	}

	idAsUUID, err := utilities.ValidateID(c.Params("userID"))
	if err != nil {
		return err
	}

	if *idAsUUID == user.ID {
		locals.SetUser(c, &user)
		return c.Next()
	}

	return utilities.Forbidden()
}
