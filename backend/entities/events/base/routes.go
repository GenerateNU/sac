package base

import (
	"github.com/GenerateNU/sac/backend/entities/events/series"
	"github.com/GenerateNU/sac/backend/entities/events/tags"

	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/garrettladley/fiberpaginate"
)

func EventRoutes(eventParams types.RouteParams) {
	eventParams.Router = EventRouter(eventParams)

	series.EventSeries(eventParams)
	tags.EventTags(eventParams)
}

func EventRouter(eventParams types.RouteParams) fiber.Router {
	eventController := NewEventController(NewEventService(eventParams.ServiceParams))

	// api/v1/events/*
	events := eventParams.Router.Group("/events")

	events.Get("/", fiberpaginate.New(), eventController.GetAllEvents)
	events.Post(
		"/",
		middleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		eventController.CreateEvent,
	)

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Patch("/", eventController.UpdateEvent)
	eventID.Delete("/", eventController.DeleteEvent)

	return events
}
