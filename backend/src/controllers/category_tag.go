package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryTagController struct {
	categoryTagService services.CategoryTagServiceInterface
}

func NewCategoryTagController(categoryTagService services.CategoryTagServiceInterface) *CategoryTagController {
	return &CategoryTagController{categoryTagService: categoryTagService}
}

func (ct *CategoryTagController) GetTagsByCategory(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	tags, err := ct.categoryTagService.GetTagsByCategory(c.Params("categoryID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tags)
}

func (ct *CategoryTagController) GetTagByCategory(c *fiber.Ctx) error {
	tag, err := ct.categoryTagService.GetTagByCategory(c.Params("categoryID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}
