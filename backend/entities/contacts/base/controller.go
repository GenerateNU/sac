package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/middleware"
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
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/contacts/{contactID}/  [get]
func (co *ContactController) GetContact(c *fiber.Ctx) error {
	contact, err := co.contactService.GetContact(c.Params("contactID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(contact)
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
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/contacts/  [get]
func (co *ContactController) GetContacts(c *fiber.Ctx) error {
	pagination, err := middleware.GetPagination(c)
	if err != nil {
		return err
	}

	contacts, err := co.contactService.GetContacts(*pagination)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(contacts)
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
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/contacts/{contactID}/  [delete]
func (co *ContactController) DeleteContact(c *fiber.Ctx) error {
	err := co.contactService.DeleteContact(c.Params("contactID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
