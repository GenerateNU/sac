package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubEvent(clubParams types.RouteParams) {
	clubEventController := controllers.NewClubEventController(services.NewClubEventService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/events/*
	events := clubParams.Router.Group("/events")

	events.Get("/", clubEventController.GetClubEvents)
}
