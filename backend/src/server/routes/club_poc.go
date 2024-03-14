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
	
	clubPointOfContact := clubIDRouter.Group("/poc")

	// api/v1/clubs/:clubID/poc/*
	clubPointOfContact.Get("/", clubPointOfContactController.GetClubPointOfContacts)
	clubPointOfContact.Get("/:pocID", clubPointOfContactController.GetClubPointOfContact)
	clubPointOfContact.Post("/", authMiddleware.ClubAuthorizeById, clubPointOfContactController.CreateClubPointOfContact)
	// clubPointOfContacts.Put("/", authMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContact)
	clubPointOfContact.Delete("/:pocID", authMiddleware.ClubAuthorizeById, clubPointOfContactController.DeleteClubPointOfContact)
}