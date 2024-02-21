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

// GetClubContacts godoc
//
// @Summary		Retrieve all contacts for a club
// @Description	Retrieves all contacts associated with a club
// @ID			get-contacts-by-club
// @Tags      	club-contact
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.Contact
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/club/{clubID}/contacts  [get]
func (cc *ClubContactController) GetClubContacts(c *fiber.Ctx) error {
	contacts, err := cc.clubContactService.GetClubContacts(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// PostContact godoc
//
// @Summary		Creates a contact
// @Description	Creates a contact
// @ID			create-contact
// @Tags      	club-contact
// @Accept		json
// @Produce		json
// @Param		clubID		path	string	true	"Club ID"
// @Param		contactBody	body	models.PutContactRequestBody	true	"Contact Body"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/club/{clubID}/contacts  [post]
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
