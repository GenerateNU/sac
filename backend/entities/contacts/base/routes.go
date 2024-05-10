package base

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func Contact(contactParams types.RouteParams) {
	contactController := NewContactController(NewContactService(contactParams.ServiceParams))

	// api/v1/contacts/*
	contacts := contactParams.Router.Group("/contacts")

	contacts.Get("/", utilities.Make(contactController.GetContacts))
	contacts.Get("/:contactID", utilities.Make(contactController.GetContact))
	contacts.Delete("/:contactID", contactParams.AuthMiddleware.Authorize(auth.DeleteAll), utilities.Make(contactController.DeleteContact))
}
