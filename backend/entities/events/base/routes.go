package base

import (
	"github.com/GenerateNU/sac/backend/entities/events/previews"
	"github.com/GenerateNU/sac/backend/entities/events/series"
	"github.com/GenerateNU/sac/backend/entities/events/tags"
	"github.com/gofiber/fiber/v2"

	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func EventRoutes(eventParams types.RouteParams) {
	eventParams.Router = EventRouter(eventParams)

	previews.EventPreviews(eventParams)
	series.EventSeries(eventParams)
	tags.EventTags(eventParams)
}

func EventRouter(eventParams types.RouteParams) fiber.Router {
	eventController := NewEventController(NewEventService(eventParams.ServiceParams))

	// api/v1/events/*
	events := eventParams.Router.Group("/events")

	events.Get("/", eventParams.UtilityMiddleware.Paginator, eventController.GetEvents)
	events.Post(
		"/",
		authMiddleware.AttachExtractor(eventParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		eventController.CreateEvent,
	)

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Patch("/", eventController.UpdateEvent)
	eventID.Delete("/", eventController.DeleteEvent)

	return events
}
