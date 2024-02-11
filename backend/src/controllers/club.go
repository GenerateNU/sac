package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubController struct {
	clubService services.ClubServiceInterface
}

func NewClubController(clubService services.ClubServiceInterface) *ClubController {
	return &ClubController{clubService: clubService}
}

func (cl *ClubController) GetAllClubs(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := cl.clubService.GetClubs(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (cl *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := cl.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

func (cl *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := cl.clubService.GetClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

func (cl *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := cl.clubService.UpdateClub(c.Params("clubID"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

func (cl *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := cl.clubService.DeleteClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Point of Contact

// UpsertPointofContact godoc
//
// @Summary		Creates or Updates a PointofContact
// @Description	Creates or Updates a PointofContact
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
		return errors.FailedToParseRequestBody.FiberError(c)
	}
	pointOfContact, err := u.clubService.UpsertPointOfContact(c.Params("clubID"), pointOfContactBody)
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
	clubId := c.Params("clubID")
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
	clubId := c.Params("clubID")
	pocId := c.Params("pocID")
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
	clubId := c.Params("clubID")
	pocId := c.Params("pocID")
	err := u.clubService.DeletePointOfContact(pocId, clubId)
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
