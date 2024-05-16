package event

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/garrettladley/fiberpaginate"
)

func ClubEvent(clubParams types.RouteParams) {
	clubEventController := NewClubEventController(NewClubEventService(clubParams.ServiceParams))

	// api/v1/clubs/:clubID/events/*
	events := clubParams.Router.Group("/events")

	events.Get("/", fiberpaginate.New(), clubEventController.GetClubEvents)
}
