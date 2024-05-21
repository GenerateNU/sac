package pocs

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubPointOfContact(clubParams types.RouteParams) {
	clubPointOfContactController := NewClubPointOfContactController(NewClubPointOfContactService(clubParams.ServiceParams))

	clubPointOfContacts := clubParams.Router.Group("/pocs")

	// api/v1/clubs/:clubID/pocs/*
	clubPointOfContacts.Get("/", clubPointOfContactController.GetClubPointOfContacts)
	clubPointOfContacts.Get("/:pocID", clubPointOfContactController.GetClubPointOfContact)
	clubPointOfContacts.Post(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubPointOfContactController.CreateClubPointOfContact,
	)
	clubPointOfContacts.Patch(
		"/:pocID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubPointOfContactController.UpdateClubPointOfContact,
	)
	clubPointOfContacts.Patch(
		"/:pocID/photo",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubPointOfContactController.UpdateClubPointOfContactPhoto,
	)
	clubPointOfContacts.Delete(
		"/:pocID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubPointOfContactController.DeleteClubPointOfContact,
	)
}
