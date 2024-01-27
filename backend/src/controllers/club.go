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

func (l *ClubController) GetAllClubs(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := l.clubService.GetClubs(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (l *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := l.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

func (l *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := l.clubService.GetClub(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

func (l *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := l.clubService.UpdateClub(c.Params("id"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

func (l *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := l.clubService.DeleteClub(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *ClubController) GetContacts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	contacts, err := l.clubService.GetContacts(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func (l *ClubController) GetClubContacts(c *fiber.Ctx) error {
	contacts, err := l.clubService.GetClubContacts(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func (l *ClubController) PutContact(c *fiber.Ctx) error {
	var contactBody models.PutContactRequestBody

	if err := c.BodyParser(&contactBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	contact, err := l.clubService.PutContact(c.Params("id"), contactBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

func (l *ClubController) DeleteContact(c *fiber.Ctx) error {
	err := l.clubService.DeleteContact(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
