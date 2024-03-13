package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type PointOfContactController struct {
	pointOfContactService services.PointOfContactServiceInterface
}

func NewPointOfContactController(pointOfContactService services.PointOfContactServiceInterface) *PointOfContactController {
	return &PointOfContactController{pointOfContactService: pointOfContactService}
}

// GetPointOfContact godoc
//
// @Summary		Retrieves a point of contact
// @Description	Retrieves a point of contact by id
// @ID			get-point-of-contact
// @Tags      	point of contact
// @Accept		json
// @Produce		json
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		201	  {object}	  models.PointOfContact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/poc/{pocID}/  [get]
func (poc *PointOfContactController) GetPointOfContact(c *fiber.Ctx) error {
	pointOfContact, err := poc.pointOfContactService.GetPointOfContact(c.Params("pocID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(pointOfContact)
}

// GetPointOfContacts godoc
//
// @Summary		Retrieve all point of contacts
// @Description	Retrieves all point of contacts
// @ID			get-point-of-contacts
// @Tags      	point of contact
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	  []models.PointOfContact
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/poc/  [get]
func (poc *PointOfContactController) GetPointOfContacts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	pointOfContacts, err := poc.pointOfContactService.GetPointOfContacts(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(pointOfContacts)
}

// DeletePointOfContact godoc
//
// @Summary		Deletes a point of contact
// @Description	Deletes a point of contact
// @ID			delete-point-of-contact
// @Tags      	point of contact
// @Accept		json
// @Produce		json
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		204	  {string}	  string
// @Failure     400   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/poc/{pocID}/  [delete]
func (poc *PointOfContactController) DeletePointOfContact(c *fiber.Ctx) error {
	err := poc.pointOfContactService.DeletePointOfContact(c.Params("pocID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

