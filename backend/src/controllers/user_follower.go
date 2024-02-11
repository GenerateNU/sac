package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type UserFollowerController struct {
	userFollowerService services.UserFollowerServiceInterface
}

func NewUserFollowerController(userFollowerService services.UserFollowerServiceInterface) *UserFollowerController {
	return &UserFollowerController{userFollowerService: userFollowerService}
}

func (uf *UserFollowerController) CreateFollowing(c *fiber.Ctx) error {
	err := uf.userFollowerService.CreateFollowing(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (uf *UserFollowerController) DeleteFollowing(c *fiber.Ctx) error {
	err := uf.userFollowerService.DeleteFollowing(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (uf *UserFollowerController) GetAllFollowing(c *fiber.Ctx) error {
	clubs, err := uf.userFollowerService.GetFollowing(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}
