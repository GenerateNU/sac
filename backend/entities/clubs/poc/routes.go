package poc

import (
	"github.com/GenerateNU/sac/backend/middleware"
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
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubPointOfContactController.CreateClubPointOfContact,
	)
	clubPointOfContacts.Patch(
		"/:pocID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubPointOfContactController.UpdateClubPointOfContact,
	)
	clubPointOfContacts.Patch(
		"/:pocID/photo",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubPointOfContactController.UpdateClubPointOfContactPhoto,
	)
	clubPointOfContacts.Delete(
		"/:pocID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubPointOfContactController.DeleteClubPointOfContact,
	)
}
