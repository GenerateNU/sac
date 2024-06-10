package auth

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) UserAuthorizeById(c *fiber.Ctx) error {
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
		return c.Next()
	}

	idAsUUID, err := utilities.ValidateID(c.Params("userID"))
	if err != nil {
		return err
	}

	if idAsUUID == &user.ID {
		return c.Next()
	}

	return utilities.Forbidden()
}
