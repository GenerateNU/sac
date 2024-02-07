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

// GetAllClubs returns all clubs.
//
// @Summary     Get all clubs
// @Description Get all clubs with pagination support
// @ID           get-all-clubs
// @Tags         club
// @Produce      json
// @Param        limit     query    int    false    "Number of clubs per page"
// @Param        page      query    int    false    "Page number"
// @Success      200       {object} []models.Club
// @Failure      400       {string} string "failed to get clubs"
// @Router       /api/v1/clubs/ [get]
func (l *ClubController) GetAllClubs(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := l.clubService.GetClubs(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

// CreateClub creates a new club.
//
// @Summary     Create a club
// @Description Create a new club
// @ID           create-club
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        clubBody     body    models.CreateClubRequestBody    true    "Club details"
// @Success      201       {object} models.Club
// @Failure      400       {string} string "failed to create club"
// @Router       /api/v1/clubs/ [post]
func (l *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := l.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

// GetClub returns a club by ID.
//
// @Summary     Get a club by ID
// @Description Get a club by its ID
// @ID           get-club
// @Tags         club
// @Produce      json
// @Param        id     path    string    true    "Club ID"
// @Success      200       {object} models.Club
// @Failure      400       {string} string "failed to get club"
// @Router       /api/v1/clubs/{id} [get]
func (l *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := l.clubService.GetClub(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

// UpdateClub updates a club by ID.
//
// @Summary     Update a club by ID
// @Description Update a club by its ID
// @ID           update-club
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        id     path    string    true    "Club ID"
// @Param        clubBody     body    models.UpdateClubRequestBody    true    "Club details"
// @Success      200       {object} models.Club
// @Failure      400       {string} string "failed to update club"
// @Router       /api/v1/clubs/{id} [put]
func (l *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := l.clubService.UpdateClub(c.Params("id"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

// DeleteClub deletes a club by ID.
//
// @Summary     Delete a club by ID
// @Description Delete a club by its ID
// @ID           delete-club
// @Tags         club
// @Produce      json
// @Param        id     path    string    true    "Club ID"
// @Success      204       -    No Content
// @Failure      400       {string} string "failed to delete club"
// @Router       /api/v1/clubs/{id} [delete]
func (l *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := l.clubService.DeleteClub(c.Params("id"))
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