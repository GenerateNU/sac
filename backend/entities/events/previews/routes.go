package previews

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GenerateNU/sac/backend/types"
)

func EventPreviews(eventParams types.RouteParams) {
	eventController := NewEventController(NewService(eventParams.ServiceParams))

	// api/v1/events/previews/
	eventParams.Router.Route("/previews", func(events fiber.Router) {
		events.Get("/", eventParams.UtilityMiddleware.Paginator, eventController.GetEventPreviews)
	})
}
