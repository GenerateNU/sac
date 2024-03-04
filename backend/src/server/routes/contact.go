package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func Contact(contactParams types.RouteParams) {
	contactController := controllers.NewContactController(services.NewContactService(contactParams.ServiceParams))

	// api/v1/contacts/*
	contacts := contactParams.Router.Group("/contacts")

	contacts.Get("/", contactController.GetContacts)
	contacts.Get("/:contactID", contactController.GetContact)
	contacts.Delete("/:contactID", contactParams.AuthMiddleware.UserAuthorizeById, contactController.DeleteContact)
}
