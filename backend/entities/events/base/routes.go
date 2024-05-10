package base

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func Event(eventParams types.RouteParams) {
	eventController := NewEventController(NewEventService(eventParams.ServiceParams))

	// api/v1/events/*
	events := eventParams.Router.Group("/events")

	events.Get("/", utilities.Make(eventController.GetAllEvents))
	events.Post("/", eventParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(eventController.CreateEvent))

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", utilities.Make(eventController.GetEvent))
	eventID.Get("/series", utilities.Make(eventController.GetSeriesByEventID))
	eventID.Patch("/", utilities.Make(eventController.UpdateEvent))
	eventID.Patch("/series", utilities.Make(eventController.UpdateSeriesByEventID))
	eventID.Delete("/", utilities.Make(eventController.DeleteEvent))
	eventID.Delete("/series", utilities.Make(eventController.DeleteSeriesByEventID))

	eventID.Get("/hosts", utilities.Make(eventController.GetHostsByEventID))
	eventID.Get("/tags", utilities.Make(eventController.GetTagsByEventID))

	// api/v1/events/series/*
	series := events.Group("/series")

	// api/v1/events/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", utilities.Make(eventController.GetSeriesByID))
	seriesID.Patch("/", eventParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(eventController.UpdateSeriesByID))
	seriesID.Delete("/", eventParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(eventController.DeleteSeriesByID))
}
