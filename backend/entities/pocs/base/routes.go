package base

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func PointOfContact(pointOfContactParams types.RouteParams) {
	pointOfContactController := NewPointOfContactController(NewPointOfContactService(pointOfContactParams.ServiceParams))

	// api/v1/pocs/*
	pointofContact := pointOfContactParams.Router.Group("/pocs")

	pointofContact.Get("/", utilities.Make(pointOfContactController.GetPointOfContacts))
	pointofContact.Get("/:pocID", utilities.Make(pointOfContactController.GetPointOfContact))
	// pointOfContact.Get("/:pocID/file", utilities.Make(pointOfContactController.GetPointOfContacFileInfo))
}
