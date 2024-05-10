package contact

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubContact(clubParams types.RouteParams) {
	clubContactController := NewClubContactController(NewClubContactService(clubParams.ServiceParams))

	clubContacts := clubParams.Router.Group("/contacts")

	// api/v1/clubs/:clubID/contacts/*
	clubContacts.Get("/", utilities.Make(clubContactController.GetClubContacts))
	clubContacts.Put("/", clubParams.AuthMiddleware.ClubAuthorizeById, utilities.Make(clubContactController.PutContact))
}
