package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func PointOfContact(router fiber.Router, pointOfContactService services.PointOfContactServiceInterface) {
	pointOfContactController := controllers.NewPointOfContactController(pointOfContactService)

	// api/v1/poc/*
	pointofContact := router.Group("/poc")

	pointofContact.Get("/", pointOfContactController.GetPointOfContacts)
	pointofContact.Get("/:pocID", pointOfContactController.GetPointOfContact)
	pointofContact.Delete("/:pocID", pointOfContactController.DeletePointOfContact)
}
