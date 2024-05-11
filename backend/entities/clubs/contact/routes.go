package contact

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/GenerateNU/sac/backend/middleware"
)

func ClubContact(clubParams types.RouteParams) {
	clubContactController := NewClubContactController(NewClubContactService(clubParams.ServiceParams))

	clubContacts := clubParams.Router.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", utilities.Make(clubContactController.GetClubContacts))
	clubContacts.Put(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubContactController.PutContact),
	)
}
