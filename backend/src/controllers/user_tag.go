package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type UserTagController struct {
	userTagService services.UserTagServiceInterface
}

func NewUserTagController(userTagService services.UserTagServiceInterface) *UserTagController {
	return &UserTagController{userTagService: userTagService}
}

func (u *UserTagController) GetUserTags(c *fiber.Ctx) error {
	tags, err := u.userTagService.GetUserTags(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(&tags)
}

func (u *UserTagController) CreateUserTags(c *fiber.Ctx) error {
	var requestBody models.CreateUserTagsBody
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tags, err := u.userTagService.CreateUserTags(c.Params("userID"), requestBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(&tags)
}
