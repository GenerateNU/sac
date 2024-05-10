package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/gofiber/fiber/v2"
)

type PointOfContactController struct {
	pointOfContactService PointOfContactServiceInterface
}

func NewPointOfContactController(pointOfContactService PointOfContactServiceInterface) *PointOfContactController {
	return &PointOfContactController{pointOfContactService: pointOfContactService}
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
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/pocs/  [get]
func (poc *PointOfContactController) GetPointOfContacts(c *fiber.Ctx) error {
	pointOfContacts, err := poc.pointOfContactService.GetPointOfContacts(c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContacts)
}

// GetPointOfContact godoc
//
// @Summary		Retrieves a point of contact
// @Description	Retrieves a point of contact by id
// @ID			get-point-of-contact
// @Tags      	point of contact
// @Produce		json
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		200	  {object}	  models.PointOfContact
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/pocs/{pocID}/  [get]
func (poc *PointOfContactController) GetPointOfContact(c *fiber.Ctx) error {
	pointOfContact, err := poc.pointOfContactService.GetPointOfContact(c.Params("pocID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContact)
}
