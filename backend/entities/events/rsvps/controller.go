package rsvps

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (co *Controller) GetEventRSVPs(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	followers, err := co.service.GetEventRSVPs(c.Params("eventID"), *pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}

func (co *Controller) CreateEventRSVP(c *fiber.Ctx) error {
	if err := co.service.CreateEventRSVP(c.Params("eventID"), c.Params("userID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusCreated)
}

func (co *Controller) DeleteEventRSVP(c *fiber.Ctx) error {
	if err := co.service.DeleteEventRSVP(c.Params("eventID"), c.Params("userID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
