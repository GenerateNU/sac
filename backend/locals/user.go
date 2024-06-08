package locals

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func UserFrom(c *fiber.Ctx) (*models.User, error) {
	iUser := c.Locals(userKey)
	if iUser == nil {
		return nil, utilities.Forbidden()
	}

	user, ok := iUser.(*models.User)
	if !ok {
		return nil, fmt.Errorf("user is not of type models.User. got: %T", user)
	}

	return user, nil
}

func SetUser(c *fiber.Ctx, user *models.User) {
	c.Locals(userKey, user)
}
