package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Event(router fiber.Router, eventService services.EventServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	eventController := controllers.NewEventController(eventService)

	// api/v1/events/*
	events := router.Group("/events")

	events.Get("/", eventController.GetAllEvents)
	events.Post("/", authMiddleware.ClubAuthorizeById, eventController.CreateEvent)

	// api/v1/events/:eventID/*
	eventID := events.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Get("/series", eventController.GetSeriesByEventID)
	eventID.Patch("/", eventController.UpdateEvent)
	eventID.Patch("/series", eventController.UpdateSeriesByEventID)
	eventID.Delete("/", eventController.DeleteEvent)
	eventID.Delete("/series", eventController.DeleteSeriesByEventID)

	// api/v1/events/series/*
	series := events.Group("/series")

	// api/v1/events/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", eventController.GetSeriesByID)
	seriesID.Patch("/", authMiddleware.ClubAuthorizeById, eventController.UpdateSeriesByID)
	seriesID.Delete("/", authMiddleware.ClubAuthorizeById, eventController.DeleteSeriesByID)
}
