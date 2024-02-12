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

func (ut *UserTagController) GetUserTags(c *fiber.Ctx) error {
	tags, err := ut.userTagService.GetUserTags(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(&tags)
}

func (ut *UserTagController) CreateUserTags(c *fiber.Ctx) error {
	var requestBody models.CreateUserTagsBody
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tags, err := ut.userTagService.CreateUserTags(c.Params("userID"), requestBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(&tags)
}
