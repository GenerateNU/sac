package rsvps

import "github.com/gofiber/fiber/v2"

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (co *Controller) CreateEventRSVP(c *fiber.Ctx) error {
	return nil
}

func (co *Controller) GetEventRSVPs(c *fiber.Ctx) error {
	return nil
}

func (co *Controller) DeleteEventRSVP(c *fiber.Ctx) error {
	return nil
}
