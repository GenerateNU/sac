package event

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubEvent(clubParams types.RouteParams) {
	clubEventController := NewClubEventController(NewClubEventService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/events/*
	events := clubParams.Router.Group("/events")

	events.Get("/", utilities.Make(clubEventController.GetClubEvents))
}
