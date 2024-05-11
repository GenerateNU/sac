package poc

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubPointOfContact(clubParams types.RouteParams) {
	clubPointOfContactController := NewClubPointOfContactController(NewClubPointOfContactService(clubParams.ServiceParams))

	clubPointOfContacts := clubParams.Router.Group("/pocs")

	// api/v1/clubs/:clubID/pocs/*
	clubPointOfContacts.Get("/", utilities.Make(clubPointOfContactController.GetClubPointOfContacts))
	clubPointOfContacts.Get("/:pocID", utilities.Make(clubPointOfContactController.GetClubPointOfContact))
	clubPointOfContacts.Post(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubPointOfContactController.CreateClubPointOfContact),
	)
	clubPointOfContacts.Patch(
		"/:pocID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubPointOfContactController.UpdateClubPointOfContact),
	)
	clubPointOfContacts.Patch(
		"/:pocID/photo",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubPointOfContactController.UpdateClubPointOfContactPhoto),
	)
	clubPointOfContacts.Delete(
		"/:pocID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubPointOfContactController.DeleteClubPointOfContact),
	)
}
