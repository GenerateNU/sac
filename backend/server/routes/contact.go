package routes

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func Contact(contactParams types.RouteParams) {
	contactController := controllers.NewContactController(services.NewContactService(contactParams.ServiceParams))

	// api/v1/contacts/*
	contacts := contactParams.Router.Group("/contacts")

	contacts.Get("/", contactController.GetContacts)
	contacts.Get("/:contactID", contactController.GetContact)
	contacts.Delete("/:contactID", contactParams.AuthMiddleware.Authorize(auth.DeleteAll), contactController.DeleteContact)
}
