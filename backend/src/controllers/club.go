package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubController struct {
	clubService services.ClubServiceInterface
}

func NewClubController(clubService services.ClubServiceInterface) *ClubController {
	return &ClubController{clubService: clubService}
}

func (cc *ClubController) GetAllClubs(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := cc.clubService.GetClubs(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (cc *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := cc.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

func (cc *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := cc.clubService.GetClub(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

func (cc *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := cc.clubService.UpdateClub(c.Params("id"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

func (cc *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := cc.clubService.DeleteClub(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
