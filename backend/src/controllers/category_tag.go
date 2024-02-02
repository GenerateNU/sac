package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryTagController struct {
	categoryTagService services.CategoryTagServiceInterface
}

func NewCategoryTagController(categoryTagService services.CategoryTagServiceInterface) *CategoryTagController {
	return &CategoryTagController{categoryTagService: categoryTagService}
}

func (t *CategoryTagController) GetTagByCategory(c *fiber.Ctx) error {
	tag, err := t.categoryTagService.GetTagByCategory(c.Params("categoryID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}
