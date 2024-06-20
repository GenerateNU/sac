package rsvps

import (
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func EventsRSVPs(params types.RouteParams) {
	controller := NewController(NewHandler(params.ServiceParams))

	// api/v1/events/:eventID/rsvps/*
	params.Router.Route("/:eventID/rsvps", func(r fiber.Router) {
		r.Get("/", params.UtilityMiddleware.Paginator, controller.GetEventRSVPs)
		r.Post("/:userID", controller.CreateEventRSVP)
		r.Delete("/:userID", params.AuthMiddleware.UserAuthorizeById, params.AuthMiddleware.Authorize(permission.DeleteAll), controller.DeleteEventRSVP)
	})
}
