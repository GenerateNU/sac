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

// CreateContact godoc
//
// @Summary		Creates a contact
// @Description	Creates a contact
// @ID			create-contact
// @Tags      	contact
// @Accept		json
// @Produce		json
// @Param		contactBody	body	models.Contact	true	"Contact Body"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/contact/  [post]
func (co *ContactController) GetContact(c *fiber.Ctx) error {
	contact, err := co.contactService.GetContact(c.Params("contactID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// GetContacts godoc
//
// @Summary		Retrieve all contacts
// @Description	Retrieves all contacts
// @ID			get-contacts
// @Tags      	contact
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	  []models.Contact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/contact/  [get]
func (co *ContactController) GetContacts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	contacts, err := co.contactService.GetContacts(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// CreateContact godoc
//
// @Summary		Creates a contact
// @Description	Creates a contact
// @ID			create-contact
// @Tags      	contact
// @Accept		json
// @Produce		json
// @Param		contactBody	body	models.Contact	true	"Contact Body"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/contact/  [post]
func (co *ContactController) DeleteContact(c *fiber.Ctx) error {
	err := co.contactService.DeleteContact(c.Params("contactID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
