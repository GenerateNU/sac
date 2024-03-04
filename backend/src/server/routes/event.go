package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func Event(eventParams types.RouteParams) {
	eventController := controllers.NewEventController(services.NewEventService(eventParams.ServiceParams))

	// api/v1/events/*
	events := eventParams.Router.Group("/events")

	events.Get("/", eventController.GetAllEvents)
	events.Post("/", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.CreateEvent)

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Get("/series", eventController.GetSeriesByEventID)
	events.Patch("/", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.UpdateEvent)
	eventID.Delete("/", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.DeleteEvent)
	eventID.Delete("/series", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.DeleteSeriesByEventID)

	// api/v1/events/:eventID/series/*
	series := events.Group("/series")

	series.Get("/", eventController.GetSeriesByEventID)
	series.Delete("/", eventController.DeleteSeriesByEventID)

	series.Get("/", eventController.GetSeriesByEventID)
	series.Delete("/", eventController.DeleteSeriesByEventID)

	// api/v1/events/:eventID/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", eventController.GetSeriesByID)
	seriesID.Patch("/", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.UpdateSeriesByID)
	seriesID.Delete("/", eventParams.AuthMiddleware.ClubAuthorizeById, eventController.DeleteSeriesByID)
}
