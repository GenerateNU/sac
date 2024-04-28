package poc

import (
	"github.com/GenerateNU/sac/backend/types"
)

func ClubPointOfContact(clubParams types.RouteParams) {
	clubPointOfContactController := NewClubPointOfContactController(NewClubPointOfContactService(clubParams.ServiceParams))

	clubPointOfContacts := clubParams.Router.Group("/pocs")

	// api/v1/clubs/:clubID/pocs/*
	clubPointOfContacts.Get("/", clubPointOfContactController.GetClubPointOfContacts)
	clubPointOfContacts.Get("/:pocID", clubPointOfContactController.GetClubPointOfContact)
	clubPointOfContacts.Post("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.CreateClubPointOfContact)
	clubPointOfContacts.Patch("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContact)
	clubPointOfContacts.Patch("/:pocID/photo", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.UpdateClubPointOfContactPhoto)
	clubPointOfContacts.Delete("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, clubPointOfContactController.DeleteClubPointOfContact)
}
