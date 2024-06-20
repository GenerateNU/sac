package rsvps

import (
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/types"
)

func EventsRSVPs(params types.RouteParams) {
	controller := NewController(NewHandler(params.ServiceParams))

	// api/v1/events/:eventID/rsvps/*
	eventRSVPS := params.Router.Group("/:eventID/rsvps")

	eventRSVPS.Get("/", params.UtilityMiddleware.Paginator, controller.GetEventRSVPs)
	eventRSVPS.Post("/:userID", params.AuthMiddleware.UserAuthorizeById, controller.CreateEventRSVP)
	eventRSVPS.Delete("/:userID", params.AuthMiddleware.UserAuthorizeById, params.AuthMiddleware.Authorize(permission.DeleteAll), controller.DeleteEventRSVP)
}
