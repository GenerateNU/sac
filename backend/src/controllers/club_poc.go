package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubPointOfContactController struct {
	clubPointOfContactService services.ClubPointOfContactServiceInterface
}

func NewClubPointOfContactController(clubPointOfContactService services.ClubPointOfContactServiceInterface) *ClubPointOfContactController {
	return &ClubPointOfContactController{clubPointOfContactService: clubPointOfContactService}
}


// func (u *ClubPointOfContactController) UpsertClubPointOfContact(c *fiber.Ctx) error {
// 	var pointOfContactBody models.CreatePointOfContactBody

// 	if err := c.BodyParser(&pointOfContactBody); err != nil {
// 		return errors.FailedToParseRequestBody.FiberError(c)
// 	}

// 	pointOfContact, err := u.clubPointOfContactService.UpsertPointOfContact(c.Params("clubID"), pointOfContactBody)
// 	if err != nil {
// 		return err.FiberError(c)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(pointOfContact)
// }

// GetClubPointOfContacts godoc
//
// @Summary		Retrieve all point of contacts for a club
// @Description	Retrieves all point of contacts associated with a club
// @ID			get-point-of-contacts-by-club
// @Tags      	club-point-of-contact
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.PointOfContact
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/poc/  [get]
func (u *ClubPointOfContactController) GetClubPointOfContacts(c *fiber.Ctx) error {
	pointOfContact, err := u.clubPointOfContactService.GetClubPointOfContacts(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(pointOfContact)
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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/poc/{pocID}  [get]
func (u *ClubPointOfContactController) GetClubPointOfContact(c *fiber.Ctx) error {
	pointOfContact, err := u.clubPointOfContactService.GetClubPointOfContact(c.Params("clubID"), c.Params("pocID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(pointOfContact)
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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/poc/{pocID}  [delete]
func (u *ClubPointOfContactController) DeleteClubPointOfContact(c *fiber.Ctx) error {
	err := u.clubPointOfContactService.DeleteClubPointOfContact(c.Params("clubID"), c.Params("pocID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
