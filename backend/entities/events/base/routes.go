package base

import (
	"github.com/GenerateNU/sac/backend/entities/events/series"
	"github.com/GenerateNU/sac/backend/entities/events/tags"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
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

	events.Get("/", utilities.Make(eventController.GetAllEvents))
	events.Post("/", eventParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(eventController.CreateEvent))

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", utilities.Make(eventController.GetEvent))
	eventID.Patch("/", utilities.Make(eventController.UpdateEvent))
	eventID.Delete("/", utilities.Make(eventController.DeleteEvent))

	return events
}
