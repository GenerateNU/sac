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
// @Failure     400   {string}    string "failed to validate point of contact"
// @Failure     400   {string}    string "failed to validate club id"
// @Failure     500   {string}    string "failed to map response to model"
// @Failure     500   {string}    string "failed to upsert point of contact"
// @Router		api/v1/clubs/:id/poc/:pocId  [put]

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


// GetAllPointOfContact godoc
//
// @Summary		Gets all point of contact
// @Description	Returns all point of contact
// @ID			get-all-point-of-contact
// @Tags      	club
// @Produce		json
// @Param		"Club ID"
// @Success		200	  {object}	  []models.PointOfContact
// @Failure     404   {string}    string "point of contact not found"
// @Failure		400   {string}    string "failed to validate point of contact"
// @Failure     500   {string}    string "failed to get point of contact"
// @Router		api/v1/clubs/:id/poc/:pocId [get]

func (u *ClubController) GetAllPointOfContact(c *fiber.Ctx) error {
	clubId := c.Params("id")
	pointOfContact, err := u.clubService.GetAllPointOfContacts(clubId)
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
// @Success		200	  {object}	  models.PointOfContact
// @Failure     404   {string}    string "point of contact not found"
// @Failure		400   {string}    string "failed to validate point of contact"
// @Failure     500   {string}    string "failed to get point of contact"
// @Router		api/v1/clubs/:id/poc/:pocId [get]

func (u *ClubController) GetPointOfContact(c *fiber.Ctx) error {
	clubId := c.Params("id")
	pocId := c.Params("pocId")
	pointOfContact, err := u.clubService.GetPointOfContact(pocId, clubId)
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(pointOfContact)
}


// DeletePointOfContact godoc
//
// @Summary		Deletes the given clubID and pocId
// @Description	Returns nil
// @ID			delete-point-of-contact
// @Tags      	club
// @Produce		json
// @Success		204    {string}        {no content}
// @Failure     500    {string}        {fail to delete point of contact}
// @Router		api/v1/clubs/:id/poc/:pocId [delete]

func (u *ClubController) DeletePointOfContact(c *fiber.Ctx) error {
	clubId := c.Params("id")
	pocId := c.Params("pocId")
	err := u.clubService.DeletePointOfContact(pocId, clubId)
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
