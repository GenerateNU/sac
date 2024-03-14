package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubContact(clubsIDRouter fiber.Router, clubContactService services.ClubContactServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	clubContactController := controllers.NewClubContactController(clubContactService)

	clubContacts := clubsIDRouter.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", clubContactController.GetClubContacts)
	clubContacts.Put("/", authMiddleware.ClubAuthorizeById, clubContactController.PutContact)
}
