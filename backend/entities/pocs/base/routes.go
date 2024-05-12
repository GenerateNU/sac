package base

import (
	"github.com/GenerateNU/sac/backend/types"
)

func PointOfContact(pointOfContactParams types.RouteParams) {
	pointOfContactController := NewPointOfContactController(NewPointOfContactService(pointOfContactParams.ServiceParams))

	// api/v1/pocs/*
	pointofContact := pointOfContactParams.Router.Group("/pocs")

	pointofContact.Get("/", pointOfContactController.GetPointOfContacts)
	pointofContact.Get("/:pocID", pointOfContactController.GetPointOfContact)
	// pointOfContact.Get("/:pocID/file", pointOfContactController.GetPointOfContacFileInfo))
}
