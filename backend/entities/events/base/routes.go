package base

import (
	"github.com/GenerateNU/sac/backend/entities/events/previews"
	"github.com/GenerateNU/sac/backend/entities/events/rsvps"
	"github.com/GenerateNU/sac/backend/entities/events/series"
	"github.com/GenerateNU/sac/backend/entities/events/tags"

	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func EventRoutes(eventParams types.RouteParams) {
	// api/v1/events/*
	events := eventParams.Router.Group("/events")
	eventParams.Router = events

	// MARK: must be called first to avoid conflict between api/v1/events/preview and api/v1/events/:eventID
	previews.EventPreviews(eventParams)
	EventRouter(eventParams)
	rsvps.EventsRSVPs(eventParams)
	series.EventSeries(eventParams)
	tags.EventTags(eventParams)
}

func EventRouter(eventParams types.RouteParams) {
	eventController := NewEventController(NewEventService(eventParams.ServiceParams))

	eventParams.Router.Get("/", eventParams.UtilityMiddleware.Paginator, eventController.GetEvents)
	eventParams.Router.Post(
		"/",
		authMiddleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		eventController.CreateEvent,
	)

	// api/v1/events/:eventID/*
	eventID := eventParams.Router.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Patch("/", eventController.UpdateEvent)
	eventID.Delete("/", eventController.DeleteEvent)
}
