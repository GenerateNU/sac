package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type LeaderController struct {
	leaderService LeaderServiceInterface
}

func NewLeaderController(leaderService LeaderServiceInterface) *LeaderController {
	return &LeaderController{leaderService: leaderService}
}

// GetLeaders godoc
//
// @Summary		Retrieve all point of contacts
// @Description	Retrieves all point of contacts
// @ID			get-point-of-contacts
// @Tags      	point of contact
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	  []models.Leader
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/leader/  [get]
func (l *LeaderController) GetLeadership(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	Leaders, err := l.leaderService.GetLeaders(*pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(Leaders)
}

// GetLeader godoc
//
// @Summary		Retrieves a point of contact
// @Description	Retrieves a point of contact by id
// @ID			get-point-of-contact
// @Tags      	point of contact
// @Produce		json
// @Param		leaderID	path	string	true	"Point of Contact ID"
// @Success		200	  {object}	  models.Leader
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/leader/{leaderID}/  [get]
func (l *LeaderController) GetLeader(c *fiber.Ctx) error {
	Leader, err := l.leaderService.GetLeader(c.Params("leaderID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(Leader)
}
