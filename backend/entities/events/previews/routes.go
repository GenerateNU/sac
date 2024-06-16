package previews

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GenerateNU/sac/backend/types"
)

func EventPreviews(eventParams types.RouteParams) {
	controller := NewController(NewService(eventParams.ServiceParams))

	// api/v1/events/previews/*
	eventParams.Router.Route("/previews", func(r fiber.Router) {
		r.Get("/", eventParams.UtilityMiddleware.Paginator, controller.GetEventPreviews)
	})
}
