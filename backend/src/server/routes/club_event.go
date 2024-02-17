package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubEvent(clubIDRouter fiber.Router, clubEventService services.ClubEventServiceInterface, middlewareService middleware.MiddlewareInterface) {
	clubEventController := controllers.NewClubEventController(clubEventService)

	// api/v1/clubs/:clubID/events/*
	events := clubIDRouter.Group("/events")

	events.Get("/", clubEventController.GetClubEvents)
}
