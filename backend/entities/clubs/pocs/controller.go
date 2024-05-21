package pocs

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type ClubPointOfContactController struct {
	clubPointOfContactService ClubPointOfContactServiceInterface
}

func NewClubPointOfContactController(clubPointOfContactService ClubPointOfContactServiceInterface) *ClubPointOfContactController {
	return &ClubPointOfContactController{clubPointOfContactService: clubPointOfContactService}
}

// GetClubPointOfContacts godoc
//
// @Summary		Retrieve all point of contacts for a club
// @Description	Retrieves all point of contacts associated with a club
// @ID			get-point-of-contacts-by-club
// @Tags      	club-point-of-contact
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.PointOfContact
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/pocs/  [get]
func (cpoc *ClubPointOfContactController) GetClubPointOfContacts(c *fiber.Ctx) error {
	pointOfContact, err := cpoc.clubPointOfContactService.GetClubPointOfContacts(c.Params("clubID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContact)
}

// GetClubPointOfContact godoc
//
// @Summary		Retrieve a point of contact for a club
// @Description	Retrieves a point of contact associated with a club
// @ID			get-point-of-contact-by-club
// @Tags      	club-point-of-contact
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		200	  {object}	    models.PointOfContact
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/pocs/{pocID}  [get]
func (cpoc *ClubPointOfContactController) GetClubPointOfContact(c *fiber.Ctx) error {
	pointOfContact, err := cpoc.clubPointOfContactService.GetClubPointOfContact(c.Params("clubID"), c.Params("pocID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContact)
}

// UpdateClubPointOfContactPhoto godoc
//
// @Summary		Update a point of contact photo for a club
// @Description	Updates a point of contact photo associated with a club
// @ID			update-point-of-contact-photo-by-club
// @Tags      	club-point-of-contact
// @Accept		multipart/form-data
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		200	  {object}	    models.PointOfContact
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/poc/{pocID}  [patch]
func (cpoc *ClubPointOfContactController) UpdateClubPointOfContactPhoto(c *fiber.Ctx) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	pointOfContact, err := cpoc.clubPointOfContactService.UpdateClubPointOfContactPhoto(c.Params("clubID"), c.Params("pocID"), formFile)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContact)
}

// UpdateClubPointOfContact godoc
//
// @Summary		Update a point of contact for a club
// @Description	Updates a point of contact associated with a club
// @ID			update-point-of-contact-by-club
// @Tags      	club-point-of-contact
// @Accept      json
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		200	  {object}	    models.PointOfContact
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/poc/{pocID}  [put]
func (cpoc *ClubPointOfContactController) UpdateClubPointOfContact(c *fiber.Ctx) error {
	var pointOfContactBody UpdatePointOfContactBody

	if err := c.BodyParser(&pointOfContactBody); err != nil {
		return utilities.InvalidJSON()
	}

	pointOfContact, err := cpoc.clubPointOfContactService.UpdateClubPointOfContact(c.Params("clubID"), c.Params("pocID"), pointOfContactBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(pointOfContact)
}

// CreateClubPointOfContact godoc
//
// @Summary		Create a point of contact for a club
// @Description	Creates a point of contact associated with a club
// @ID			create-point-of-contact-by-club
// @Tags      	club-point-of-contact
// @Accept		multipart/form-data
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		201	  {object}	    models.PointOfContact
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/poc/  [post]
func (cpoc *ClubPointOfContactController) CreateClubPointOfContact(c *fiber.Ctx) error {
	var pointOfContactBody CreatePointOfContactBody

	if err := c.BodyParser(&pointOfContactBody); err != nil {
		return utilities.InvalidJSON()
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	pointOfContact, err := cpoc.clubPointOfContactService.CreateClubPointOfContact(c.Params("clubID"), pointOfContactBody, formFile)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(pointOfContact)
}

// DeleteClubPointOfContact godoc
//
// @Summary		Delete a point of contact for a club
// @Description	Delete a point of contact associated with a club
// @ID			delete-point-of-contact-by-club
// @Tags      	club-point-of-contact
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		pocID	path	string	true	"Point of Contact ID"
// @Success		204	  {object}	    nil
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/poc/{pocID}  [delete]
func (cpoc *ClubPointOfContactController) DeleteClubPointOfContact(c *fiber.Ctx) error {
	err := cpoc.clubPointOfContactService.DeleteClubPointOfContact(c.Params("clubID"), c.Params("pocID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
