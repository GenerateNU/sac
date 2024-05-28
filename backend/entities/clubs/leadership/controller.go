package leadership

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type ClubLeaderController struct {
	clubLeaderService ClubLeaderServiceInterface
}

func NewClubLeaderController(clubLeaderService ClubLeaderServiceInterface) *ClubLeaderController {
	return &ClubLeaderController{clubLeaderService: clubLeaderService}
}

// GetClubLeadership godoc
//
// @Summary		Retrieve all leadership for a club
// @Description	Retrieves all leadership associated with a club
// @ID			get-leadership-by-club
// @Tags      	club-leader
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.Leader
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/  [get]
func (l *ClubLeaderController) GetClubLeadership(c *fiber.Ctx) error {
	leadership, err := l.clubLeaderService.GetClubLeadership(c.Params("clubID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(leadership)
}

// GetClubLeader godoc
//
// @Summary		Retrieve a leader for a club
// @Description	Retrieves a leader associated with a club
// @ID			get-leader-by-club
// @Tags      	club-leader
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		leaderID	path	string	true	"leader ID"
// @Success		200	  {object}	    models.Leader
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/{leaderID}  [get]
func (l *ClubLeaderController) GetClubLeader(c *fiber.Ctx) error {
	leader, err := l.clubLeaderService.GetClubLeader(c.Params("clubID"), c.Params("leaderID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(leader)
}

// UpdateClubLeaderPhoto godoc
//
// @Summary		Update a leader photo for a club
// @Description	Updates a leader photo associated with a club
// @ID			update-leader-photo-by-club
// @Tags      	club-leader
// @Accept		multipart/form-data
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		leaderID	path	string	true	"leader ID"
// @Success		200	  {object}	    models.Leader
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/{leaderID}  [patch]
func (l *ClubLeaderController) UpdateClubLeaderPhoto(c *fiber.Ctx) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	leader, err := l.clubLeaderService.UpdateClubLeaderPhoto(c.Params("clubID"), c.Params("leaderID"), formFile)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(leader)
}

// UpdateClubLeader godoc
//
// @Summary		Update a leader for a club
// @Description	Updates a leader associated with a club
// @ID			update-leader-by-club
// @Tags      	club-leader
// @Accept      json
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		leaderID	path	string	true	"leader ID"
// @Success		200	  {object}	    models.Leader
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/{leaderID}  [put]
func (l *ClubLeaderController) UpdateClubLeader(c *fiber.Ctx) error {
	var leaderBody UpdateLeaderBody

	if err := c.BodyParser(&leaderBody); err != nil {
		return utilities.InvalidJSON()
	}

	leader, err := l.clubLeaderService.UpdateClubLeader(c.Params("clubID"), c.Params("leaderID"), leaderBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(leader)
}

// CreateClubLeader godoc
//
// @Summary		Create a leader for a club
// @Description	Creates a leader associated with a club
// @ID			create-leader-by-club
// @Tags      	club-leader
// @Accept		multipart/form-data
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		201	  {object}	    models.Leader
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/  [post]
func (l *ClubLeaderController) CreateClubLeader(c *fiber.Ctx) error {
	var leaderBody CreateLeaderBody

	if err := c.BodyParser(&leaderBody); err != nil {
		return utilities.InvalidJSON()
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	leader, err := l.clubLeaderService.CreateClubLeader(c.Params("clubID"), leaderBody, formFile)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(leader)
}

// DeleteClubLeader godoc
//
// @Summary		Delete a leader for a club
// @Description	Delete a leader associated with a club
// @ID			delete-leader-by-club
// @Tags      	club-leader
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		leaderID	path	string	true	"leader ID"
// @Success		204	  {object}	    nil
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/leadership/{leaderID}  [delete]
func (l *ClubLeaderController) DeleteClubLeader(c *fiber.Ctx) error {
	err := l.clubLeaderService.DeleteClubLeader(c.Params("clubID"), c.Params("leaderID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
