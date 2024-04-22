package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubContact(clubParams types.RouteParams) {
	clubContactController := controllers.NewClubContactController(services.NewClubContactService(clubParams.ServiceParams))

	clubContacts := clubParams.Router.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", clubContactController.GetClubContacts)
	clubContacts.Put("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubContactController.PutContact)
}
