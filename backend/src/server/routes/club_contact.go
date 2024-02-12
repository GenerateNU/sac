package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubContact(clubsIDRouter fiber.Router, clubContactService services.ClubContactServiceInterface) {
	clubContactController := controllers.NewClubContactController(clubContactService)

	clubContacts := clubsIDRouter.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", clubContactController.GetClubContacts)
	clubContacts.Put("/", clubContactController.PutContact)
}
