package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Contact(router fiber.Router, contactService services.ContactServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	contactController := controllers.NewContactController(contactService)

	// api/v1/contacts/*
	contacts := router.Group("/contacts")

	contacts.Get("/", contactController.GetContacts)
	contacts.Get("/:contactID", contactController.GetContact)
	contacts.Delete("/:contactID", authMiddleware.UserAuthorizeById, contactController.DeleteContact)
}
