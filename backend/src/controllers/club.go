package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

// Point of Contact
type ClubController struct {
	clubService services.ClubServiceInterface
}

func NewClubController(clubService services.ClubServiceInterface) *ClubController {
	return &ClubController{clubService: clubService}
}

// UpsertPointofContact godoc
//
// @Summary		Creates or Updates a User
// @Description	Creates or Updates a User
// @ID			upsert-point-of-contact
// @Tags      	club
// @Accept      json
// @Produce		json
// @Success		201	  {object}	  models.PointOfContact
// @Failure     400   {string}    string "failed to create or update point of contact"
// @Failure     500   {string}    string "internal server error"
// @Router		api/v1/clubs/:id/poc/:email  [put]
func (u *ClubController) UpsertPointOfContact(c *fiber.Ctx) error {
	var pointOfContactBody models.CreatePointOfContactBody
	if err := c.BodyParser(&pointOfContactBody); err != nil {
		return errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToParseRequestBody}.FiberError(c)
	}
	pointOfContact, err := u.clubService.UpsertPointOfContact(c.Params("id"), pointOfContactBody)
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(pointOfContact)
}

// GetPointOfContact godoc
//
// @Summary		Gets a user
// @Description	Returns a user
// @ID			get-point-of-contact
// @Tags      	club
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Success		200	  {object}	  models.PointOfContact
// @Failure     404   {string}    string "point of contact not found"
// @Failure		400   {string}    string "failed to validate point of contact"
// @Failure     500   {string}    string "failed to get point of contact"
// @Router		api/v1/clubs/:id/poc/:email [get]
func (u *ClubController) GetAllPointOfContact(c *fiber.Ctx) error {
	clubId := c.Params("id")
	pointOfContact, err := u.clubService.GetAllPointOfContacts(clubId)
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(pointOfContact)
}

// DeletePointOfContact godoc
//
// @Summary		Deletes the given clubID and userID
// @Description	Returns nil
// @ID			delete-point-of-contact
// @Tags      	club
// @Produce		json
// @Success		204    {string}        {no content}
// @Failure     500    {string}        {fail to delete point of contact}
// @Router		api/v1/clubs/:id/poc/:email [delete]

func (u *ClubController) DeletePointOfContact(c *fiber.Ctx) error {
	clubId := c.Params("id")
	email := c.Params("email")
	err := u.clubService.DeletePointOfContact(email, clubId)
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
