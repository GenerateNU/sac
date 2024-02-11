package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubContactController struct {
	clubContactService services.ClubContactServiceInterface
}

func NewClubContactController(clubContactService services.ClubContactServiceInterface) *ClubContactController {
	return &ClubContactController{clubContactService: clubContactService}
}

func (cc *ClubContactController) GetClubContacts(c *fiber.Ctx) error {
	contacts, err := cc.clubContactService.GetClubContacts(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func (cc *ClubContactController) PutContact(c *fiber.Ctx) error {
	var contactBody models.PutContactRequestBody

	if err := c.BodyParser(&contactBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	contact, err := cc.clubContactService.PutClubContact(c.Params("clubID"), contactBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}
