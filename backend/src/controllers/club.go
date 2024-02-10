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

// GetContact returns a contact by ID.
//
// @Summary     Get a contact by ID
// @Description Get a contact by its ID
// @ID           get-contact
// @Tags         club
// @Produce      json
// @Param        id         path    string    true    "Contact ID"
// @Success      200       {object} models.Contact
// @Failure      400       {string} string "failed to get contact"
// @Router       /api/v1/contacts/{id} [get]
func (l *ClubController) GetContact(c *fiber.Ctx) error {
	contact, err := l.clubService.GetContact(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// GetContacts returns all contacts.
//
// @Summary     Get all contacts
// @Description Get all contacts with pagination support
// @ID           get-all-contacts
// @Tags         club
// @Produce      json
// @Param        limit     query    int    false    "Number of contacts per page"
// @Param        page      query    int    false    "Page number"
// @Success      200       {object} []models.Contact
// @Failure      400       {string} string "failed to get contacts"
// @Router       /api/v1/contacts/ [get]
func (l *ClubController) GetContacts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	contacts, err := l.clubService.GetContacts(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// GetClubContacts returns all contacts of a club.
//
// @Summary     Get all contacts of a club
// @Description Get all contacts of a club with pagination support
// @ID           get-club-contacts
// @Tags         club
// @Produce      json
// @Param        id     path    string    true    "Club ID"
// @Success      200       {object} []models.Contact
// @Failure      400       {string} string "failed to get club contacts"
// @Router       /api/v1/clubs/{id}/contacts [get]
func (l *ClubController) GetClubContacts(c *fiber.Ctx) error {
	contacts, err := l.clubService.GetClubContacts(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// PutContact creates or updates a contact for a club.
//
// @Summary     Create or update a contact for a club
// @Description Creates a contact for a club if it does not exist, otherwise updates an existing contact
// @ID           put-contact
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        id     path    string    true    "Club ID"
// @Param        contactBody     body    models.PutContactRequestBody    true    "Contact details"
// @Success      200       {object} models.Contact
// @Failure      400       {string} string "failed to create/update contact"
// @Router       /api/v1/clubs/{id}/contacts [put]
func (l *ClubController) PutContact(c *fiber.Ctx) error {
	var contactBody models.PutContactRequestBody

	if err := c.BodyParser(&contactBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	contact, err := l.clubService.PutContact(c.Params("id"), contactBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// DeleteContact deletes a contact by ID.
//
// @Summary     Delete a contact by ID
// @Description Delete a contact by its ID
// @ID           delete-contact
// @Tags         club
// @Produce      json
// @Param        id     path    string    true    "Contact ID"
// @Success      204       -    No Content
// @Failure      400       {string} string "failed to delete contact"
// @Router       /api/v1/contacts/{id} [delete]
func (l *ClubController) DeleteContact(c *fiber.Ctx) error {
	err := l.clubService.DeleteContact(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}