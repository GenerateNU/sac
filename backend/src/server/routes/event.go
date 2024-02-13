package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Event(router fiber.Router, eventService services.EventServiceInterface) {
	eventController := controllers.NewEventController(eventService)

	events := router.Group("/events")

	events.Get("/:id", eventController.GetEvent)
	events.Get("/:id/series", eventController.GetSeriesByEventId)
	events.Get("/", eventController.GetAllEvents)
	events.Post("/", eventController.CreateEvent)
	events.Patch("/:id", eventController.UpdateEvent)
	events.Delete("/:id", eventController.DeleteEvent)
	events.Delete("/:id/series", eventController.DeleteEventSeries) 
}