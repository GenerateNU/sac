package contact

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type ClubContactController struct {
	clubContactService ClubContactServiceInterface
}

func NewClubContactController(clubContactService ClubContactServiceInterface) *ClubContactController {
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/contacts/  [get]
func (cc *ClubContactController) GetClubContacts(c *fiber.Ctx) error {
	contacts, err := cc.clubContactService.GetClubContacts(c.Params("clubID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(contacts)
}

// PutContact godoc
//
// @Summary		Creates a contact
// @Description	Creates a contact
// @ID			put-contact
// @Tags      	club-contact
// @Accept		json
// @Produce		json
// @Param		clubID		path	string	true	"Club ID"
// @Param		contactBody	body	models.PutContactRequestBody	true	"Contact Body"
// @Success		201	  {object}	  models.Contact
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/clubs/{clubID}/contacts/  [put]
func (cc *ClubContactController) PutContact(c *fiber.Ctx) error {
	var contactBody models.PutContactRequestBody

	if err := c.BodyParser(&contactBody); err != nil {
		return utilities.InvalidJSON()
	}

	contact, err := cc.clubContactService.PutClubContact(c.Params("clubID"), contactBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(contact)
}
