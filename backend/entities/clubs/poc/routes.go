package poc

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubPointOfContact(clubParams types.RouteParams) {
	clubPointOfContactController := NewClubPointOfContactController(NewClubPointOfContactService(clubParams.ServiceParams))

	clubPointOfContacts := clubParams.Router.Group("/pocs")

	// api/v1/clubs/:clubID/pocs/*
	clubPointOfContacts.Get("/", utilities.Make(clubPointOfContactController.GetClubPointOfContacts))
	clubPointOfContacts.Get("/:pocID", utilities.Make(clubPointOfContactController.GetClubPointOfContact))
	clubPointOfContacts.Post("/", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubPointOfContactController.CreateClubPointOfContact))
	clubPointOfContacts.Patch("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubPointOfContactController.UpdateClubPointOfContact))
	clubPointOfContacts.Patch("/:pocID/photo", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubPointOfContactController.UpdateClubPointOfContactPhoto))
	clubPointOfContacts.Delete("/:pocID", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubPointOfContactController.DeleteClubPointOfContact))
}
