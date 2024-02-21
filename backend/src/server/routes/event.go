package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Event(router fiber.Router, eventService services.EventServiceInterface) {
	eventController := controllers.NewEventController(eventService)

	// api/v1/events/*
	events := router.Group("/events")

	events.Get("/", eventController.GetAllEvents)
	events.Post("/", eventController.CreateEvent)
	events.Patch("/", eventController.UpdateEvent)

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Delete("/", eventController.DeleteEvent)

	// api/v1/events/:eventID/series/*
	series := router.Group("/series")

	series.Get("/", eventController.GetSeriesByEventID)
	series.Delete("/", eventController.DeleteSeriesByEventID)

	// api/v1/events/:eventID/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", eventController.GetSeriesByID)
	seriesID.Patch("/", eventController.UpdateSeriesByID)
	seriesID.Delete("/", eventController.DeleteSeriesByID)
}
