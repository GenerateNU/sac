package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubTagController struct {
	clubTagService services.ClubTagServiceInterface
}

func NewClubTagController(clubTagService services.ClubTagServiceInterface) *ClubTagController {
	return &ClubTagController{clubTagService: clubTagService}
}

// CreateClubTags godoc
//
// @Summary		Create Club Tags
// @Description	Adds Tags to a Club 
// @ID			create-club-tags
// @Tags      	club
// @Accept      json
// @Produce		json
// @Success		201	  {object}	  []models.Tag
// @Failure     404   {string}    string "club not found"
// @Failure 	400   {string}    string "invalid request body"
// @Failure		400   {string}    string "failed to validate id"
// @Failure		500   {string}	  string "database error"
// @Router		/api/v1/clubs/:id/tags  [post]
func (l *ClubTagController) CreateClubTags(c *fiber.Ctx) error {
	var clubTagsBody models.CreateClubTagsRequestBody
	if err := c.BodyParser(&clubTagsBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	clubTags, err := l.clubTagService.CreateClubTags(c.Params("id"), clubTagsBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(clubTags)
}

// GetClubTags godoc
//
// @Summary		Get Club Tags
// @Description	Retrieves the tags for a club
// @ID			get-club-tags
// @Tags      	club
// @Produce		json
// @Success		200	  {object}	  []models.Tag
// @Failure     404   {string}    string "club not found"
// @Failure 	400   {string}    string "invalid request body"
// @Failure		400   {string}    string "failed to validate id"
// @Failure		500   {string}	  string "database error"
// @Router		/api/v1/clubs/:id/tags  [get]
func (l *ClubTagController) GetClubTags(c *fiber.Ctx) error {
	clubTags, err := l.clubTagService.GetClubTags(c.Params("id"))

	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubTags)
}

// DeleteClubTags godoc
//
// @Summary		Delete Club Tags
// @Description	Deletes the tags for a club
// @ID			delete-club-tags
// @Tags      	club
// @Success		204	  
// @Failure     404   {string}    string "club not found"
// @Failure 	400   {string}    string "invalid request body"
// @Failure		400   {string}    string "failed to validate id"
// @Failure		500   {string}	  string "database error"
// @Router		/api/v1/clubs/:id/tags/:tagId  [delete]
func (l *ClubTagController) DeleteClubTag(c *fiber.Ctx) error {
	err := l.clubTagService.DeleteClubTag(c.Params("id"), c.Params("tagId"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)

}