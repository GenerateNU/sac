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
	event := router.Group("/events")

	event.Get("/", eventController.GetAllEvents)
	event.Post("/", authMiddleware.ClubAuthorizeById, eventController.CreateEvent)

	// api/v1/events/:eventID/*
	eventID := event.Group("/:eventID")

	eventID.Get("/", eventController.GetEvent)
	eventID.Patch("/", authMiddleware.ClubAuthorizeById, eventController.UpdateEvent)
	eventID.Delete("/", authMiddleware.ClubAuthorizeById, eventController.DeleteEvent)

	// api/v1/events/:eventID/series/*
	series := event.Group("/series")

	series.Get("/", eventController.GetSeriesByEventID)
	series.Delete("/", authMiddleware.ClubAuthorizeById, eventController.DeleteSeriesByEventID)

	// api/v1/events/:eventID/series/:seriesID/*
	seriesID := series.Group("/:seriesID")

	seriesID.Get("/", eventController.GetSeriesByID)
	seriesID.Patch("/", authMiddleware.ClubAuthorizeById, eventController.UpdateSeriesByID)
	seriesID.Delete("/", authMiddleware.ClubAuthorizeById, eventController.DeleteSeriesByID)
}
