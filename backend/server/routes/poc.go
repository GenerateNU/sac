package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func PointOfContact(pointOfContactParams types.RouteParams) {
	pointOfContactController := controllers.NewPointOfContactController(services.NewPointOfContactService(pointOfContactParams.ServiceParams))

	// api/v1/pocs/*
	pointofContact := pointOfContactParams.Router.Group("/pocs")

	pointofContact.Get("/", pointOfContactController.GetPointOfContacts)
	pointofContact.Get("/:pocID", pointOfContactController.GetPointOfContact)
	// pointOfContact.Get("/:pocID/file", pointOfContactController.GetPointOfContacFileInfo)
}
