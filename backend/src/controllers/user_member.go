package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type UserMemberController struct {
	clubMemberService services.UserMemberServiceInterface
}

func NewUserMemberController(clubMemberService services.UserMemberServiceInterface) *UserMemberController {
	return &UserMemberController{clubMemberService: clubMemberService}
}

func (um *UserMemberController) CreateMembership(c *fiber.Ctx) error {
	err := um.clubMemberService.CreateMembership(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (um *UserMemberController) DeleteMembership(c *fiber.Ctx) error {
	err := um.clubMemberService.DeleteMembership(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (um *UserMemberController) GetMembership(c *fiber.Ctx) error {
	followers, err := um.clubMemberService.GetMembership(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
