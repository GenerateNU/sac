package contacts

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubContact(clubParams types.RouteParams) {
	clubContactController := NewClubContactController(NewClubContactService(clubParams.ServiceParams))

	clubContacts := clubParams.Router.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", clubContactController.GetClubContacts)
	clubContacts.Put(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubContactController.PutContact,
	)
}
