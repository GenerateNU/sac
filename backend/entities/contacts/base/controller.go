package base

import (
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/gofiber/fiber/v2"
)

type ContactController struct {
	contactService ContactServiceInterface
}

func NewContactController(contactService ContactServiceInterface) *ContactController {
	return &ContactController{contactService: contactService}
}

// GetContact godoc
//
// @Summary		Retrieves a contact
// @Description	Retrieves a contact by id
// @ID			get-contact
// @Tags      	contact
// @Accept		json
// @Produce		json
// @Param		contactID	path	string	true	"Contact ID"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/contacts/{contactID}/  [get]
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
// @Router		/contacts/  [get]
func (co *ContactController) GetContacts(c *fiber.Ctx) error {
	contacts, err := co.contactService.GetContacts(c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// DeleteContact godoc
//
// @Summary		Deletes a contact
// @Description	Deletes a contact
// @ID			delete-contact
// @Tags      	contact
// @Accept		json
// @Produce		json
// @Param		contactID	path	string	true	"Contact ID"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/contacts/{contactID}/  [delete]
func (co *ContactController) DeleteContact(c *fiber.Ctx) error {
	err := co.contactService.DeleteContact(c.Params("contactID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
