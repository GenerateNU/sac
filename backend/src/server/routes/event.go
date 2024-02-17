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

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Get("/series", eventController.GetSeriesByEventID)
	events.Patch("/", eventController.UpdateEvent)
	eventID.Delete("/", eventController.DeleteEvent)
	eventID.Delete("/series", eventController.DeleteSeriesByEventID)

	// api/v1/events/:eventID/series/*
	series := router.Group("/series")

	// api/v1/events/:eventID/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", eventController.GetSeriesByID)
	seriesID.Patch("/", eventController.UpdateSeriesByID)
	seriesID.Delete("/", eventController.DeleteSeriesByID)
}
