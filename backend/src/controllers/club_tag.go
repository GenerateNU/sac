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
// @Summary		Create club tags
// @Description	Creates tags for a club
// @ID			create-club-tags
// @Tags      	club-tag
// @Accept		json
// @Produce		json
// @Param		clubID		path	string	true	"Club ID"
// @Param		clubTagsBody	body	models.CreateClubTagsRequestBody	true	"Club Tags Body"
// @Success		201	  {object}	  []models.Tag
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/club/{clubID}/tags  [post]
func (l *ClubTagController) CreateClubTags(c *fiber.Ctx) error {
	var clubTagsBody models.CreateClubTagsRequestBody
	if err := c.BodyParser(&clubTagsBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	clubTags, err := l.clubTagService.CreateClubTags(c.Params("clubID"), clubTagsBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(clubTags)
}

// GetClubTags godoc
//
// @Summary		Retrieve all tags for a club
// @Description	Retrieves all tags associated with a club
// @ID			get-tags-by-club
// @Tags      	club-tag
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/club/{clubID}/tags  [get]
func (l *ClubTagController) GetClubTags(c *fiber.Ctx) error {
	clubTags, err := l.clubTagService.GetClubTags(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubTags)
}

// DeleteClubTag godoc
//
// @Summary		Delete a tag for a club
// @Description	Deletes a tag associated with a club
// @ID			delete-tag-by-club
// @Tags      	club-tag
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		tagID	path	string	true	"Tag ID"
// @Success		204	  {string}	    utilites.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/club/{clubID}/tags/{tagID}  [delete]
func (l *ClubTagController) DeleteClubTag(c *fiber.Ctx) error {
	err := l.clubTagService.DeleteClubTag(c.Params("clubID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
