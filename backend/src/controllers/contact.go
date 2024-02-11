package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ContactController struct {
	contactService services.ContactServiceInterface
}

func NewContactController(contactService services.ContactServiceInterface) *ContactController {
	return &ContactController{contactService: contactService}
}

func (co *ContactController) GetContact(c *fiber.Ctx) error {
	contact, err := co.contactService.GetContact(c.Params("contactID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

func (co *ContactController) GetContacts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	contacts, err := co.contactService.GetContacts(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

func (co *ContactController) DeleteContact(c *fiber.Ctx) error {
	err := co.contactService.DeleteContact(c.Params("contactID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
