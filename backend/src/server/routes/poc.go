package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func PointOfContact(router fiber.Router, pointOfContactService services.PointOfContactServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	pointOfContactController := controllers.NewPointOfContactController(pointOfContactService)

	// api/v1/pocs/*
	pointofContact := router.Group("/pocs")

	pointofContact.Get("/", pointOfContactController.GetPointOfContacts)
	pointofContact.Get("/:pocID", pointOfContactController.GetPointOfContact)
	// pointOfContact.Get("/:pocID/file", pointOfContactController.GetPointOfContacFileInfo)
}
