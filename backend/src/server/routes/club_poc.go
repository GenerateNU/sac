package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubPointOfContact(clubParams types.RouteParams) {
	clubPointOfContactController := controllers.NewClubPointOfContactController(services.NewClubPointOfContactService(clubParams.ServiceParams))

	clubPointOfContacts := clubParams.Router.Group("/pocs")

	// api/v1/clubs/:clubID/pocs/*
	clubPointOfContacts.Get("/", clubPointOfContactController.GetClubPointOfContacts)
	clubPointOfContacts.Get("/:pocID", clubPointOfContactController.GetClubPointOfContact)
	clubPointOfContacts.Post("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.CreateClubPointOfContact)
	clubPointOfContacts.Patch("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContact)
	clubPointOfContacts.Patch("/:pocID/photo", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContactPhoto)
	clubPointOfContacts.Delete("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.DeleteClubPointOfContact)
}
