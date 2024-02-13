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

func (cl *ClubController) GetAllClubs(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := cl.clubService.GetClubs(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (cl *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := cl.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

func (cl *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := cl.clubService.GetClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

func (cl *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := cl.clubService.UpdateClub(c.Params("clubID"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

func (cl *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := cl.clubService.DeleteClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}


func (l *ClubController) GetClubEvents(c *fiber.Ctx) error {
	//TODO add filters by date
	if events, err := l.clubService.GetClubEvents(c.Params("id")); err != nil {
		return err.FiberError(c)
	} else {
		return c.Status(fiber.StatusOK).JSON(events)
	}
}

// func (l *ClubController) GetClubEvents(c *fiber.Ctx) error {
// 	defaultLimit := 10
// 	defaultPage := 1

// 	events, err := l.clubService.GetClubEvents(c.Params("cid"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
// 	if err != nil {
// 		return err.FiberError(c)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(events)
// }