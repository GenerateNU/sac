package events

import (
	"github.com/GenerateNU/sac/backend/types"
)

func ClubEvent(clubParams types.RouteParams) {
	clubEventController := NewClubEventController(NewClubEventService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/events/*
	events := clubParams.Router.Group("/events")

	events.Get("/", clubParams.UtilityMiddleware.Paginator, clubEventController.GetClubEvents)
}
