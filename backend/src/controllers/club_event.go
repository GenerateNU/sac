package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubEventController struct {
	clubEventService services.ClubEventServiceInterface
}

func NewClubEventController(clubEventService services.ClubEventServiceInterface) *ClubEventController {
	return &ClubEventController{clubEventService: clubEventService}
}

func (cl *ClubEventController) GetClubEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	if events, err := cl.clubEventService.GetClubEvents(c.Params("clubID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage))); err != nil {
		return err.FiberError(c)
	} else {
		return c.Status(fiber.StatusOK).JSON(events)
	}
}
