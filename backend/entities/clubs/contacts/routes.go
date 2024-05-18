package contacts

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubContact(clubParams types.RouteParams) {
	clubContactController := NewClubContactController(NewClubContactService(clubParams.ServiceParams))

	clubContacts := clubParams.Router.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", clubContactController.GetClubContacts)
	clubContacts.Put(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubContactController.PutContact,
	)
}
