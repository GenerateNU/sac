package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

// // api/v1/clubs/:clubID/poc/*
// pointOfContact := router.Group("/clubs/:clubID/poc")
// pointOfContact.Get("/", clubController.GetAllPointOfContact)
// pointOfContact.Get("/:pocID", clubController.GetPointOfContact)
// pointOfContact.Put("/", clubController.UpsertPointOfContact)
// pointOfContact.Delete("/:pocID", clubController.DeletePointOfContact)

func ClubPointOfContact(clubIDRouter fiber.Router, clubPointOfContactService services.ClubPointOfContactServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	clubPointOfContactController := controllers.NewClubPointOfContactController(clubPointOfContactService)
	
	clubPointOfContacts := clubIDRouter.Group("/poc")

	// api/v1/clubs/:clubID/poc/*
	clubPointOfContacts.Get("/", clubPointOfContactController.GetClubPointOfContacts)
	clubPointOfContacts.Get("/:pocID", clubPointOfContactController.GetClubPointOfContact)
	// clubPointOfContacts.Put("/", authMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContact)
	clubPointOfContacts.Delete("/:pocID", authMiddleware.ClubAuthorizeById, clubPointOfContactController.DeleteClubPointOfContact)
}